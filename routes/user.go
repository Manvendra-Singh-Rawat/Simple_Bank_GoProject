package routes

import (
	"SimpleBank/models"
	"SimpleBank/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var user models.Users
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := utils.HashPassword(user.HashedPassword)
	if err != nil {
		context.JSON(http.StatusInternalServerError, "some error occured, please try again later")
		return
	}
	user.HashedPassword = hashedPassword

	err = user.CreateUser()
	if err != nil {
		// context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create new user"})
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "New user created successfully"})
}
