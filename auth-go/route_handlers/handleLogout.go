package route_handlers

import (
	"auth-go/database"
	"auth-go/jwt_auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleLogout(c *gin.Context) {

	//authHeader := c.GetHeader("Authorization")

	tokenString, err := c.Cookie("token")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization failed",
		})
		return
	}

	claims, err := jwt_auth.ValidateToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = database.DB.Exec(`
	INSERT INTO token_blacklist (token,expires_at)
	VALUES($1,$2)
	`, tokenString,
		claims.ExpiresAt.Time,
	)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.SetCookie("token", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged Out successfully",
	})

}
