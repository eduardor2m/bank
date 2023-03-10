package controllers

import (
	"fmt"
	"net/http"

	"github.com/eduardor2m/bank/database"
	"github.com/eduardor2m/bank/models"
	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	numberAccountOrigin := c.Params.ByName("account_origin")
	numberAccountDestination := c.Params.ByName("account_destination")
	valueTransaction := c.Params.ByName("value")

	fmt.Println(valueTransaction)

	var accountOrigin models.Account
	result := database.DB.Where("account_number = ?", numberAccountOrigin).First(&accountOrigin)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account origin not found"})
		return
	}

	var accountDestination models.Account
	result = database.DB.Where("account_number = ?", numberAccountDestination).First(&accountDestination)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account destination not found"})
		return
	}

	var transaction models.Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := transaction.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if accountOrigin.Balance < transaction.Value {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insufficient balance"})
		return
	}

	accountOrigin.Balance -= transaction.Value
	accountDestination.Balance += transaction.Value

	result = database.DB.Save(&accountOrigin)

	if result.Error != nil {
		panic(result.Error)
	}

	result = database.DB.Save(&accountDestination)

	if result.Error != nil {
		panic(result.Error)
	}

	transaction.NumberAccountOrigin = numberAccountOrigin
	transaction.NumberAccountDestination = numberAccountDestination
	transaction.AccountID = accountOrigin.ID

	result = database.DB.Create(&transaction)

	if result.Error != nil {
		panic(result.Error)
	}

	c.JSON(http.StatusCreated, transaction)

}

func ListTransactions(c *gin.Context) {
	var transactions []models.Transaction
	result := database.DB.Find(&transactions)

	if result.Error != nil {
		panic(result.Error)
	}

	c.JSON(http.StatusOK, transactions)
}
