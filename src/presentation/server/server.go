package server

import (
	"api/src/presentation/endpoints"
	"github.com/gin-gonic/gin"
	"log"
)

// InitServer start the web server
func InitServer() {
	router := gin.New()
	router.GET("/health_check", endpoints.HealthCheck)
	router.GET("/book", endpoints.ConsultBook)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
