package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware, JWT token doğrulamasını yapar
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authorization header'dan JWT token alınır
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// Bearer token formatını doğrula
		tokenString := strings.Split(authHeader, "Bearer ")
		if len(tokenString) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		// Token'ı doğrula
		token, err := ValidateToken(tokenString[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Token geçerli ise kullanıcı bilgilerini context'e ekle
		c.Set("user", token)
		c.Next()
	}
}
