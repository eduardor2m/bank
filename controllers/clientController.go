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

	createClient := `INSERT INTO clients (name, cpf, email, password) VALUES (?, ?, ?, ?)`

	result, err := database.DB.Exec(createClient, client.Name, client.CPF, client.Email, client.Password)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"id": id})

}

func ListClients(c *gin.Context) {
	rows, err := database.DB.Query("SELECT * FROM clients")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var clients []models.Client

	for rows.Next() {
		var client models.Client
		if err := rows.Scan(&client.ID, &client.Name, &client.CPF, &client.Email, &client.Password); err != nil {
			panic(err)
		}
		clients = append(clients, client)
	}

	c.JSON(http.StatusOK, clients)
}

func GetClient(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var client models.Client

	row := database.DB.QueryRow("SELECT * FROM clients WHERE id = ?", id)
	if err := row.Scan(&client.ID, &client.Name, &client.CPF, &client.Email, &client.Password); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}

	c.JSON(http.StatusOK, client)
}

func UpdateClient(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var client models.Client

	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := client.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateClient := `UPDATE clients SET name = ?, cpf = ?, email = ?, password = ? WHERE id = ?`

	statement, err := database.DB.Prepare(updateClient)

	if err != nil {
		panic(err)
	}

	_, err = statement.Exec(client.Name, client.CPF, client.Email, client.Password, id)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Client updated successfully!"})
}

func DeleteClient(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deleteClient := `DELETE FROM clients WHERE id = ?`

	statement, err := database.DB.Prepare(deleteClient)

	if err != nil {
		panic(err)
	}

	_, err = statement.Exec(id)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Client deleted successfully!"})

}
