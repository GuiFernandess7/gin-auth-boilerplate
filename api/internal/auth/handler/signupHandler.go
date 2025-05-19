package handler

import (
	"fmt"
	"gin-auth-boilerplate/config"
	"gin-auth-boilerplate/internal/auth/schemas"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	user schemas.User
}

func (s *SignupController) checkIfUserExists(email string) (bool, error) {
	var exists bool
	err := config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (s *SignupController) addUserToDB() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s.user.HashedPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password: %v", err)
	}

	_, err = config.DB.Exec(
		"INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3)",
		s.user.Name, s.user.Email, string(hashedPassword),
	)
	if err != nil {
		return fmt.Errorf("error inserting user into DB: %v", err)
	}
	return nil
}

func Signup(c *gin.Context) {
	var req schemas.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	signupController := SignupController{
		user: schemas.User{
			Email:          req.Email,
			HashedPassword: req.HashedPassword,
		},
	}

	exists, err := signupController.checkIfUserExists(req.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	if exists {
		c.JSON(400, gin.H{"error": "Email already registered"})
		return
	}

	if err := signupController.addUserToDB(); err != nil {
		c.JSON(500, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(201, gin.H{
		"message": "Success.",
		"user": gin.H{
			"email": req.Email,
		},
	})
}
