package main

import (
	"auth-go/database"
	"auth-go/migrations"
	"auth-go/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()
	migrations.CreateTables()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	routes.RegisterRoutes(router)

	router.Run("localhost:8000")

}
