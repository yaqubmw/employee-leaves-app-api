package middleware

import (
	"employeeleave/utils/security"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var h authHeader
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			// untuk menghentikan proses di bawahnya atau lanjutan dan proses ini akan dikembalikan dalam response http ke client
			c.Abort()
			return
		}

		tokenHeader := strings.Replace(h.AuthorizationHeader, "Bearer ", "", 1)
		if tokenHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		claims, err := security.VerifyAccessToken(tokenHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
		fmt.Println("claims:", claims["username"])
		c.Set("claims", claims)
		c.Next()
	}
}
