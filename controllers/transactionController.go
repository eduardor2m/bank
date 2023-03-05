package controllers

import (
	"fmt"
	"net/http"

	"github.com/eduardor2m/bank/database"
	"github.com/eduardor2m/bank/models"
	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	idOrigin := c.Params.ByName("idOrigin")
	idDestination := c.Params.ByName("idDestination")
	fmt.Println(idOrigin)
	fmt.Println(idDestination)
	if err := c.ShouldBindUri(&idOrigin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var accountOrigin models.Account

	row := database.DB.QueryRow("SELECT * FROM accounts WHERE id = ?", idOrigin)
	if err := row.Scan(&accountOrigin.ID, &accountOrigin.Balance, &accountOrigin.AgencyNumber, &accountOrigin.AccountNumber); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}
	var accountDestination models.Account

	row = database.DB.QueryRow("SELECT * FROM accounts WHERE id = ?", idDestination)
	if err := row.Scan(&accountDestination.ID, &accountDestination.Balance, &accountDestination.AgencyNumber, &accountDestination.AccountNumber); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
		return
	}

	updateAccountOrigin := `UPDATE accounts SET balance = ? WHERE id = ?`
	statement, err := database.DB.Prepare(updateAccountOrigin)
	if err != nil {
		panic(err)
	}
	_, err = statement.Exec(accountOrigin.Balance-transaction.Value, accountOrigin.ID)

	if err != nil {
		panic(err)
	}

	updateAccountDestination := `UPDATE accounts SET balance = ? WHERE id = ?`
	statement, err = database.DB.Prepare(updateAccountDestination)
	if err != nil {
		panic(err)
	}

	_, err = statement.Exec(accountDestination.Balance+transaction.Value, accountDestination.ID)

	if err != nil {
		panic(err)
	}

	createTransaction := `INSERT INTO transactions (account_origin_id, account_destination_id, value) VALUES (?, ?, ?)`

	result, err := database.DB.Exec(createTransaction, transaction.AccountOriginID, transaction.AccountDestinationId, transaction.Value)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"id": id})

}

func ListTransactions(c *gin.Context) {
	rows, err := database.DB.Query("SELECT * FROM transactions")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var transactions []models.Transaction

	for rows.Next() {
		var transaction models.Transaction
		if err := rows.Scan(&transaction.ID, &transaction.AccountOriginID, &transaction.AccountDestinationId, &transaction.Value); err != nil {
			panic(err)
		}
		transactions = append(transactions, transaction)
	}

	c.JSON(http.StatusOK, transactions)
}
