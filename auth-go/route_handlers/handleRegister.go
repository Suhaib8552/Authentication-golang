package route_handlers

import (
	"auth-go/database"
	"auth-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HandleRegister(c *gin.Context) {

	var req models.RegisterReq

	var err error

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var count int

	err = database.DB.QueryRow(
		`SELECT COUNT(*)
			 FROM USERS
			 WHERE username = $1 OR email = $2`,
		req.Username,
		req.Email,
	).Scan(&count)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Database error",
		})
		return
	}

	if count > 0 {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"error": "Username or Email already exists",
		})
		return
	}

	hashedpass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	var userId int

	err = database.DB.QueryRow(
		`INSERT INTO USERS(username,email,password_hash) 
			VALUES($1,$2,$3)
			RETURNING id`,
		req.Username,
		req.Email,
		string(hashedpass),
	).Scan(&userId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"userId":  userId,
	})

}
