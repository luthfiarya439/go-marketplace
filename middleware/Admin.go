package middleware

import (
	"fmt"
	"go-marketplace/config"
	"go-marketplace/models"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AdminMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		response := gin.H{
			"status":  "failed",
			"message": "Header Authorization tidak ada",
		}
		c.JSON(http.StatusUnauthorized, response)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		response := gin.H{
			"status":  "failed",
			"message": "Format token tidak valid",
		}
		c.JSON(http.StatusUnauthorized, response)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString := authToken[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token salah atau sudah kadaluarsa"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
		c.Abort()
		return
	}

	if float64(time.Now().Unix()) > claims["expired"].(float64) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token kadaluarsa"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var user models.User

	config.DB.Where("id = ?", claims["id"]).Where("role = ?", "admin").Find(&user)

	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("currentUser", user)

	c.Next()
}
