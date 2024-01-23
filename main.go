package main

import (
	"github.com/gin-gonic/gin"
	"vs45tech.com/event/db"
	"vs45tech.com/event/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
