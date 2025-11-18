package main

import (
	"vasvault/internal/repositories"
	"vasvault/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	db, err := repositories.Connect()
	if err != nil {
		panic("Failed to connect to database")
	}

	r := gin.Default()
	routes.InitRoutes(r, db.DB)

	if err := r.Run(); err != nil {
		panic("Failed to run server")
	}
}
