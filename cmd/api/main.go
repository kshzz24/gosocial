package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kshzz24/gosocial/internal/database"
	"github.com/kshzz24/gosocial/internal/handlers"
	"github.com/kshzz24/gosocial/internal/middleware"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	err = database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer database.Close()
	router := gin.New()
	router.Use(gin.Logger())

	auth := router.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
		auth.POST("/forgot-password", handlers.ForgotPassword)
		auth.POST("/reset-password", handlers.ResetPassword)
	}

	api := router.Group("/api")
	api.Use(middleware.RequireAuth())
	{
		api.GET("/me", handlers.GetMe)
		api.POST("/logout", handlers.Logout)
		api.POST("/update-password", handlers.ChangePassword)

	}

	log.Println("🚀 Server is ready!")
	log.Fatal(router.Run(":" + port))

}
