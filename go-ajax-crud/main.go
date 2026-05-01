package main

import (
	"go-ajax-crud/db"
	"go-ajax-crud/handlers"
	"go-ajax-crud/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitMongo()

	r := gin.Default()

	// ✅ STATIC FILES (IMPORTANT)
	r.Static("/static", "./static")

	// ✅ HTML TEMPLATES
	r.LoadHTMLGlob("templates/*")

	// 🌐 PAGES
	r.GET("/", handlers.ShowHome)
	r.GET("/login", handlers.ShowLogin)
	r.GET("/register", handlers.ShowRegister)

	// 🔐 AUTH
	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Register)

	// 🔒 PROTECTED API
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/books", handlers.GetBooks)
		api.POST("/books", handlers.CreateBook)
		api.PUT("/books/:id", handlers.UpdateBook)
		api.DELETE("/books/:id", handlers.DeleteBook)
	}

	r.Run(":8080")
}