package controllers

import (
	"AuthBase/internal/db"
	"AuthBase/internal/models"
	"AuthBase/internal/utils"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func MagicLinkGet(c *gin.Context) {
	type Req struct {
		Email string `json:"email"`
	}
	var req Req
	c.ShouldBindBodyWithJSON(&req) // check for error
	fmt.Println(req.Email)
	checkEmail := utils.CheckEmail(req.Email)
	if checkEmail == false { //false == something wrong with email
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something wring with email"})
		return
	}

	link, err := utils.CreateMagicLink(req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Magic link creation fail"})
		return
	}

	// log.Println("Magic link:", link)
	utils.SendEmail(req.Email, link)
	c.JSON(http.StatusOK, gin.H{"success": true, "redirect": "/emailsent", "message": "Magic link sent to your email"})

}

func VerifyMagicLinkRegister(c *gin.Context) {
	token := c.Query("token")

	val, ok := utils.MagicTokens.Load(token)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	magicTokenExtract := val.(utils.MagicToken)
	utils.MagicTokens.Delete(token)

	user := &models.User{
		Email: magicTokenExtract.UserEmail,
	}
	//deley for DB
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	exitsts := db.CheckIfExists(magicTokenExtract.UserEmail)
	if exitsts == 0 {
		err := db.CreateUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User creation failed"})
			return
		}
	}

	// Generate JWT
	jwtToken, err := utils.CreateJWT(magicTokenExtract.UserEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	utils.MagicTokens.Delete(token)
	c.SetCookie("jwt", jwtToken, 3600, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/admindashboard")
}
func EmailSent(c *gin.Context) {
	c.HTML(http.StatusOK, "email-sent.html", nil)
}
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)

}
