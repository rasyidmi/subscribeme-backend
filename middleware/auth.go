package middleware

import (
	"context"
	"net/http"
	"projects-subscribeme-backend/utils"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware : to verify all authorized operations
func AuthMiddleware(c *gin.Context) error {
	var err error
	firebaseAuth := c.MustGet("firebaseAuth").(*auth.Client)
	authorizationToken := c.GetHeader("Authorization")
	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))
	if idToken == "" {
		err = utils.TokenNotAvailable{}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Id token not available"})
		c.Abort()
		return err
	}
	//verify token
	token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		c.Abort()
		return err
	}
	c.Set("UUID", token.UID)
	c.Next()

	return err
}
