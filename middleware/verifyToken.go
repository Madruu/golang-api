package middleware

import (
	"authentication-api/utils"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

//AuthenticationMiddleware checks if the user has a valid JWT token

func AuthenticationMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization");
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authentication Token"})
			c.Abort();
			return 
		}

		tokenParts := strings.Split(tokenString, " ");
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token"});
			c.Abort();
			return;
		}

		tokenString = tokenParts[1];

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication Token"})
			c.Abort();
			return;
		}

		c.Set("user_id", claims["userId"])
		c.Next();
	}
}
