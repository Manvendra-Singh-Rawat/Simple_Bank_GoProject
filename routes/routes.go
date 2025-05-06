package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/accounts", createNewAccount)     // Working
	server.GET("/accounts/:id", getAccountDetails) // Working
	server.POST("/transfermoney", transferMoney)
	server.POST("/addmoney", addMoney) // Working
}
