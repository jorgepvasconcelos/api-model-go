package server

import (
	"api/src/presentation/endpoints"
	"github.com/gin-gonic/gin"
	"log"
)

// InitServer returns a mux.Router
func InitServer() {
	router := gin.Default()
	router.GET("/book", endpoints.ConsultBook)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
