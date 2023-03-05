package main

import (
	"fmt"

	"github.com/eduardor2m/bank/database"
	"github.com/eduardor2m/bank/routes"
)

func main() {
	routes := routes.SetupRouter()
	database.InitDB()
	routes.Run(fmt.Sprintf(":%d", 8080))
}
