package middleware

import (
	"net/http"
	"projects-subscribeme-backend/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware : to verify all authorized operations
func AuthMiddleware(c *gin.Context) error {
	var err error
	authorizationToken := c.GetHeader("Authorization")
	extractedToken := ExtractToken(authorizationToken)
	if extractedToken == "" {
		err = utils.TokenNotAvailable{}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not available."})
		c.Abort()
		return err
	}
	//verify token
	payload, err := utils.VerifyAccessToken(extractedToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid."})
		c.Abort()
		return err
	}
	c.Request.Header.Add("userId", payload.UserId)
	c.Request.Header.Add("role", payload.Role)
	c.Next()

	return err
}

func ExtractToken(token string) string {
	jwtToken := strings.TrimSpace(strings.Replace(token, "Bearer", "", 1))
	return jwtToken
}
