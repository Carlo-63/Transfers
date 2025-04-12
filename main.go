package main

import (
	"log"
	"test/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.SetupRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
