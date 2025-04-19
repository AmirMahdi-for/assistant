package helpers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ParseJSONBody(c *gin.Context) (map[string]interface{}, bool) {
	var body map[string]interface{}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, false
	}

	fmt.Println("JSON Request Body:", body)
	return body, true
}
