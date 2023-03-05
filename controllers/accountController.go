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
	if err := account.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createAccount := `INSERT INTO accounts (balance, agency, account) VALUES (?, ?, ?)`

	result, err := database.DB.Exec(createAccount, account.Balance, account.AgencyNumber, account.AccountNumber)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"id": id})

}

func ListAccounts(c *gin.Context) {
	rows, err := database.DB.Query("SELECT * FROM accounts")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var accounts []models.Account

	for rows.Next() {
		var account models.Account
		if err := rows.Scan(&account.ID, &account.Balance, &account.AgencyNumber, &account.AccountNumber); err != nil {
			panic(err)
		}
		accounts = append(accounts, account)
	}

	c.JSON(http.StatusOK, accounts)
}

func GetAccount(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var account models.Account

	row := database.DB.QueryRow("SELECT * FROM accounts WHERE id = ?", id)
	if err := row.Scan(&account.ID, &account.Balance, &account.AgencyNumber, &account.AccountNumber); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	c.JSON(http.StatusOK, account)
}

func UpdateAccount(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var account models.Account

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := account.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateAccount := `UPDATE accounts SET balance = ?, agency = ?, account = ? WHERE id = ?`

	statement, err := database.DB.Prepare(updateAccount)

	if err != nil {
		panic(err)
	}

	_, err = statement.Exec(account.Balance, account.AgencyNumber, account.AccountNumber, id)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account updated successfully!"})
}

func DeleteAccount(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deleteAccount := `DELETE FROM accounts WHERE id = ?`

	statement, err := database.DB.Prepare(deleteAccount)

	if err != nil {
		panic(err)
	}

	_, err = statement.Exec(id)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account deleted successfully!"})

}
