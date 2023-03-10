package controllers

import (
	"net/http"

	"github.com/eduardor2m/bank/database"
	"github.com/eduardor2m/bank/models"
	"github.com/gin-gonic/gin"
)

func CreateClient(c *gin.Context) {
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := client.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var clientExists models.Client

	userExists := database.DB.Where("cpf = ?", client.CPF).First(&clientExists)

	if userExists.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CPF already exists"})
		return
	}
	client.Account = models.Account{Balance: 0}

	result := database.DB.Create(&client)
	if result.Error != nil {
		panic(result.Error)
	}

	c.JSON(http.StatusCreated, client)

}

func ListClients(c *gin.Context) {
	var clients []models.Client
	result := database.DB.Model(&clients).Preload("Account").Find(&clients)

	if result.Error != nil {
		panic(result.Error)
	}

	c.JSON(http.StatusOK, clients)
}

func GetClient(c *gin.Context) {
	id := c.Params.ByName("id")

	var client models.Client
	result := database.DB.First(&client, id)

	if result.Error != nil {
		panic(result.Error)
	}

	c.JSON(http.StatusOK, client)
}

func UpdateClient(c *gin.Context) {
	id := c.Params.ByName("id")

	var client models.Client
	result := database.DB.First(&client, id)

	if result.Error != nil {
		panic(result.Error)
	}

	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result = database.DB.Save(&client)

	if result.Error != nil {
		panic(result.Error)
	}

	c.JSON(http.StatusOK, client)
}

func DeleteClient(c *gin.Context) {
	id := c.Params.ByName("id")

	var client models.Client
	result := database.DB.First(&client, id)

	if result.Error != nil {
		panic(result.Error)
	}

	result = database.DB.Delete(&client)

	if result.Error != nil {
		panic(result.Error)
	}

	c.JSON(http.StatusOK, client)

}
