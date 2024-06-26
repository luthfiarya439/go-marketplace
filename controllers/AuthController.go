package controllers

import (
	"go-marketplace/config"
	"go-marketplace/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Authenticate(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		response := gin.H{
			"status":  500,
			"message": "Error, " + err.Error(),
		}
		c.JSON(500, response)
		return
	}

	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		response := gin.H{
			"status":  500,
			"message": "Email atau password anda salah",
		}
		c.JSON(500, response)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		response := gin.H{
			"status":  500,
			"message": "Email atau password anda salah",
		}
		c.JSON(500, response)
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      user.ID,
		"expired": time.Now().Add(time.Hour * 2).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		response := gin.H{
			"status":  500,
			"message": "Gagal membuat token",
		}
		c.JSON(500, response)
		return
	}

	response := gin.H{
		"status":  200,
		"message": "Berhasil membuat token",
		"token":   token,
	}

	c.JSON(200, response)
}

func Register(c *gin.Context) {

	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		response := gin.H{
			"status":  500,
			"message": "Error, " + err.Error(),
		}
		c.JSON(500, response)
		return
	}

	var user models.User

	if err := config.DB.Model(&user).Where("email = ?", input.Email).First(&user).RowsAffected; err == 1 {
		response := gin.H{
			"status":  500,
			"message": "Email sudah digunakan",
		}
		c.JSON(500, response)
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		response := gin.H{
			"status":  500,
			"message": "Gagal",
		}
		c.JSON(500, response)
		return
	}

	user = models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(passwordHash),
		Role:     "user",
	}

	if err := config.DB.Create(&user).Error; err != nil {
		response := gin.H{
			"status":  500,
			"message": "Gagal",
		}
		c.JSON(500, response)
		return
	}

	response := gin.H{
		"status":  201,
		"message": "Berhasil membuat akun",
	}

	c.JSON(201, response)
}

func GetProfile(c *gin.Context) {
	user, _ := c.Get("currentUser")

	response := gin.H{
		"status":  "success",
		"message": "Berhasil ambil data user",
		"data":    user,
	}
	c.JSON(200, response)
}
