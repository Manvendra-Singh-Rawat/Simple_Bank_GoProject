package main

import (
	"SimpleBank/db"
	"SimpleBank/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.DBConnection()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
