package middleware

import (
	"auth-go/database"
	"auth-go/jwt_auth"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing Authorization Header",
			})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		var count int

		err := database.DB.QueryRow(
			`
			SELECT COUNT(*)
			FROM token_blacklist
			WHERE token=$1
			`, tokenString,
		).Scan(&count)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if count > 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token blacklisted",
			})
			return

		}

		claims, err := jwt_auth.ValidateToken(tokenString)

		fmt.Printf("Claims: %+v\n", claims)

		if err != nil {

			fmt.Println("JWT ERROR:", err)

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.Set("email", claims.Email)
		c.Set("user_id", claims.UserID)

		c.Next()

	}

}
