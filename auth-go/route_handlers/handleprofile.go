package route_handlers

import (
	"auth-go/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handleprofile(c *gin.Context) {

	userId := c.GetInt("user_id")

	fmt.Println("User ID from token:", userId)

	var profile struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	err := database.DB.QueryRow(
		`
		SELECT id,username,email
		FROM users
		WHERE id=$1
		`, userId,
	).Scan(&profile.ID,
		&profile.Username,
		&profile.Email)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, profile)

}
