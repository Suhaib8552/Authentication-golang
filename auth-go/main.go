package main

import (
	"auth-go/database"
	"auth-go/migrations"
	"auth-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()
	migrations.CreateTables()

	router := gin.Default()

	routes.RegisterRoutes(router)

	router.Run("localhost:8000")

}
