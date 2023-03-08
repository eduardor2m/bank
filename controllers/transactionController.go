package controllers

import (
	"fmt"
	"net/http"

	"github.com/eduardor2m/bank/database"
	"github.com/eduardor2m/bank/models"
	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	idOrigin := c.Params.ByName("id_origin")
	idDestination := c.Params.ByName("id_destination")
	valueTransaction := c.Params.ByName("value")

	fmt.Println(valueTransaction)

	var accountOrigin models.Account
	result := database.DB.First(&accountOrigin, idOrigin)

	if result.Error != nil {
		panic(result.Error)
	}

	var accountDestination models.Account
	result = database.DB.First(&accountDestination, idDestination)

	if result.Error != nil {
		panic(result.Error)
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

	transaction.AccountOriginID = int(accountOrigin.ID)
	transaction.AccountDestinationId = int(accountDestination.ID)

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
