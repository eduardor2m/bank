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

	return r
}
