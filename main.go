package main

import (
	"github.com/joaoscorissa/gin-api-rest/database"
	"github.com/joaoscorissa/gin-api-rest/routes"
)

func main() {
	database.ConnectDB()
	routes.HandleRequest()
}
