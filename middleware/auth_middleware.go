package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func CheckLogin(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			c.Abort()
			return
		}

		if strings.Contains(token, "|") {
			tokenParts := strings.Split(token, "|")
			token = tokenParts[1]
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		var count int64
		result := db.Raw("SELECT COUNT(*) FROM personal_access_tokens WHERE token = ?", token).Scan(&count)
		if result.Error != nil || count == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
