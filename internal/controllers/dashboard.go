package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminDashboard(c *gin.Context) {
	// email, _ := c.Get("email")
	c.HTML(http.StatusOK, "adminDashboard.html", nil)
}
