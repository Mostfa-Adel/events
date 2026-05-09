package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-events/db"
	"github.com/go-events/routes"
)

func main() {
	db.InitDb()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")

}
