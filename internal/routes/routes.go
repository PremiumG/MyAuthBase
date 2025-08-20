package routes

import (
	"AuthBase/internal/controllers"
	"AuthBase/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	root := r.Group("/")
	{
		root.GET("/", controllers.Index)
		root.GET("/signup", controllers.SignUp)
		root.GET("/login", controllers.Login)
		root.POST("/magicLinkGet", controllers.MagicLinkGet)
		root.GET("/verifymagicregister", controllers.VerifyMagicLinkRegister)
		root.GET("/admindashboard", middleware.AuthMiddleware(), controllers.AdminDashboard)
		root.GET("/emailsent", controllers.EmailSent)
		r.NoRoute(controllers.NoRoute) //404

	}
}

// r.LoadHTMLGlob("static/*.html")
// r.GET("/", index)
// r.GET("/register", register)
// r.POST("/registerurl", registerURL)
// //r.POST("/login", loginWithMagicLink)

// r.GET("/verify", verifyMagicLink)
// r.GET("/protected", authMiddleware(), locked)
