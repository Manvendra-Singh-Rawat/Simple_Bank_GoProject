package routes

import (
	"SimpleBank/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterRoutes(server *gin.Engine) {

	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		v.RegisterValidation("currency", utils.ValidCurrency)
	}

	server.POST("/accounts", createNewAccount)     // Working
	server.GET("/accounts/:id", getAccountDetails) // Working
	server.POST("/transfermoney", transferMoney)
	server.POST("/addmoney", addMoney) // Working
}
