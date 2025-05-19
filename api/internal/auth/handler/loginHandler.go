package handler

import (
	"database/sql"
	"fmt"
	"gin-auth-boilerplate/config"
	"gin-auth-boilerplate/internal/auth/schemas"
	"gin-auth-boilerplate/internal/auth/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	db   *sql.DB
	user schemas.User
}

func NewLoginController(db *sql.DB) *LoginController {
	return &LoginController{
		db: db,
	}
}

func (l *LoginController) GetUserByEmail(email string) error {
	var passwordHash string
	fmt.Println("Tentando fazer select...")
	err := l.db.QueryRow(
		"SELECT id, email, password_hash FROM users WHERE email = $1",
		email,
	).Scan(&l.user.ID, &l.user.Email, &passwordHash)

	fmt.Println("Select feito.")
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user not found")
		}
		return fmt.Errorf("error querying user: %v", err)
	}

	l.user.HashedPassword = passwordHash
	return nil
}

func (l *LoginController) HashPassword(req schemas.LoginRequest) error {
	err := bcrypt.CompareHashAndPassword([]byte(l.user.HashedPassword), []byte(req.Password))
	if err != nil {
		return fmt.Errorf("invalid password")
	}
	return nil
}

func (l *LoginController) GenerateToken() (string, error) {
	token, err := utils.GenerateToken(l.user.ID)
	if err != nil {
		return "", fmt.Errorf("error generating token")
	}
	return token, nil
}

func Login(c *gin.Context) {
	var req schemas.LoginRequest
	fmt.Println("Recebendo request body...")
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	loginController := NewLoginController(config.DB)

	fmt.Println("Capturando usuario...")
	if err := loginController.GetUserByEmail(req.Email); err != nil {
		c.JSON(401, gin.H{"error": fmt.Sprintf("Invalid credentials: %s", err)})
		return
	}

	fmt.Println("Hashing password...")
	if err := loginController.HashPassword(req); err != nil {
		c.JSON(401, gin.H{"error": fmt.Sprintf("Invalid credentials: %s", err)})
		return
	}

	fmt.Println("Generating token...")
	token, err := loginController.GenerateToken()
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Something went wrong: %s", err)})
		return
	}

	c.JSON(200, schemas.AuthResponse{
		Token: token,
		User: schemas.User{
			ID:    loginController.user.ID,
			Email: loginController.user.Email,
		},
	})
}
