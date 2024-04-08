package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

func verifyIDToken(ctx context.Context, idToken string) *auth.Token {
	app, err := firebase.NewApp(context.Background(), &firebase.Config{
		ProjectID: "shareef-money",
	})
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

	return token
}

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

	token := verifyIDToken(context.Background(), tokenString)
	if token == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		c.Abort()
		return
	}

	c.Next()
}
