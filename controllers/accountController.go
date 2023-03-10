package controllers

import (
	"net/http"

	"github.com/eduardor2m/bank/database"
	"github.com/eduardor2m/bank/models"
	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	var account models.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var client models.Client
	var accountExists models.Account

	result := database.DB.First(&client, account.ClientID)

	if result.Error != nil {
		panic(result.Error)
	}

	database.DB.Where("client_id = ?", account.ClientID).First(&accountExists)

	if accountExists.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Client already has an account"})
		return
	}

	if accountExists.AccountNumber == account.AccountNumber {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Account number already exists"})
		return
	}

	var accounts []models.Account

	database.DB.Find(&accounts)

	for _, rangeAccount := range accounts {

		if rangeAccount.AccountNumber == account.AccountNumber {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Account number already exists"})
			return
		}
	}

	client.Account = account

	database.DB.Save(&client)

	c.JSON(http.StatusOK, client)

}

func ListAccounts(c *gin.Context) {
	var accounts []models.Account
	result := database.DB.Find(&accounts)

	if result.Error != nil {
		panic(result.Error)
	}

	c.JSON(http.StatusOK, accounts)
}

func GetAccount(c *gin.Context) {
	id := c.Params.ByName("id")

	var account models.Account

	result := database.DB.First(&account, id)

	if result.Error != nil {
		panic(result.Error)
	}

	c.JSON(http.StatusOK, account)
}

func UpdateAccount(c *gin.Context) {
	id := c.Params.ByName("id")

	var account models.Account
	result := database.DB.First(&account, id)

	if result.Error != nil {
		panic(result.Error)
	}

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&account)

	c.JSON(http.StatusOK, account)
}

func DeleteAccount(c *gin.Context) {
	id := c.Params.ByName("id")

	var account models.Account
	result := database.DB.First(&account, id)

	if result.Error != nil {
		panic(result.Error)
	}

	database.DB.Delete(&account)

	c.JSON(http.StatusOK, account)

}
