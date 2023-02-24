package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck(c *gin.Context) {
	response := map[string]string{
		"message": "I still working",
	}
	c.SecureJSON(http.StatusOK, response)
}
