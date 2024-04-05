package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CheckToken(c *gin.Context) {
	tokenHeader := c.GetHeader("Authorization")

	if tokenHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Provide authorization token",
		})
		c.Abort()
		return
	}

	if !strings.HasPrefix(tokenHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization format"})
		c.Abort()
		return
	}

	tokenString := tokenHeader[7:]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing token %v", token.Header["alg"])
		}

		publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(os.Getenv("NEXTAUTH_SECRET")))

		if err != nil {
			return nil, err
		}

		return publicKey, nil
	})

	log.Println("From Middleware", "err", err)

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	log.Println("From Middleware", tokenHeader, tokenHeader == "")

	c.Next()
}
