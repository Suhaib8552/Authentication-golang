package route_handlers

import (
	"auth-go/database"
	"auth-go/jwt_auth"
	"auth-go/models"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HandleLogin(c *gin.Context) {

	var req models.LoginReq

	if err := c.ShouldBindJSON(&req); err != nil {

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	var (
		userId        int
		email         string
		password_hash string
	)

	err := database.DB.QueryRow(`
	SELECT id,email,password_hash
	FROM USERS
	WHERE email=$1
	`,
		req.Email,
	).Scan(
		&userId,
		&email,
		&password_hash,
	)

	if err == sql.ErrNoRows {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "database error",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(password_hash),
		[]byte(req.Password),
	)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid password",
		})
		return
	}

	fmt.Println("userId from database:", userId)
	fmt.Println("email from database:", email)

	token, err := jwt_auth.GenerateToken(userId, email)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": "Successfully generated the token",
		"token":  token,
	})

}
