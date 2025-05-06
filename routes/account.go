package routes

import (
	"SimpleBank/models"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func createNewAccount(context *gin.Context) {
	var newAccount models.Accounts
	err := context.ShouldBindBodyWithJSON(&newAccount)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = newAccount.CreateAccount()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "account created successfully"})
}

func addMoney(context *gin.Context) {
	var addMoney models.AddMoney
	err := context.ShouldBindJSON(&addMoney)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	account, err := models.GetAccount(int64(addMoney.AccountID))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "account doesn't exist or currency didn't match"})
		return
	}

	if addMoney.Currency != account.Currency {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Currency matching failed"})
		return
	}

	err = models.AddMoneyToAccount(addMoney.AccountID, addMoney.Amount)
	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "No rows exists"})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Some error occured while adding money"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Money successfully added to account"})
}

func getAccountDetails(context *gin.Context) {
	accountId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	accountDetails, err := models.GetAccount(accountId)
	if err == sql.ErrNoRows {
		context.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	} else if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, *accountDetails)
}
