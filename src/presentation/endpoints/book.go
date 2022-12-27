package endpoints

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ConsultBook insere um usuário no banco de dados
func ConsultBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})

	fmt.Println("A requisicao chegou")
}
