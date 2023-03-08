package routes

import (
	"github.com/eduardor2m/bank/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	accounts := r.Group("/accounts")
	{
		accounts.POST("/", controllers.CreateAccount)
		accounts.GET("/", controllers.ListAccounts)
		accounts.GET("/:id", controllers.GetAccount)
		accounts.PUT("/:id", controllers.UpdateAccount)
		accounts.DELETE("/:id", controllers.DeleteAccount)
	}

	clients := r.Group("/clients")
	{
		clients.POST("/", controllers.CreateClient)
		clients.GET("/", controllers.ListClients)
		clients.GET("/:id", controllers.GetClient)
		clients.PUT("/:id", controllers.UpdateClient)
		clients.DELETE("/:id", controllers.DeleteClient)
	}

	transactions := r.Group("/transactions")
	{
		transactions.POST("/:id_origin/:id_destination", controllers.CreateTransaction)
		transactions.GET("/", controllers.ListTransactions)
	}

	return r
}
