package handlers

import "github.com/gin-gonic/gin"

// 🏠 Dashboard page
func ShowHome(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

// 🔐 Login page
func ShowLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

// 🟢 Register page
func ShowRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}