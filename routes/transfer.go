package routes

import (
	"SimpleBank/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func transferMoney(context *gin.Context) {

	// Gets JSON data into object
	var transferRequest models.TransferMoney
	err := context.ShouldBindJSON(&transferRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse data"})
		return
	}

	// Validate if the sender's account and currency are valid or not
	err = models.ValidAccount(int64(transferRequest.FromAccountID), transferRequest.Currency)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// Validate if the recievers's account and currency are valid or not
	err = models.ValidAccount(int64(transferRequest.ToAccountID), transferRequest.Currency)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Create transfer request
	err = transferRequest.CreateTransfer()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "transfer complete"})
}
