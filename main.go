package main

import (
	"hrm-api/config"
	"hrm-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectMongoDB()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
