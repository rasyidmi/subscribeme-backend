package middlewares

import (
	"errors"
	"net/http"
	"projects-subscribeme-backend/dto/response"
	"projects-subscribeme-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(role string) gin.HandlerFunc {
	return func(context *gin.Context) {

		authorizationToken := context.GetHeader("Authorization")
		tokenString := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))

		if strings.Contains(authorizationToken, "Bearer") == false {
			response.Error(context, "failed", http.StatusUnauthorized, errors.New("request does not contain an access token"))
			context.Abort()
			return
		}

		if tokenString == "" {
			response.Error(context, "failed", http.StatusUnauthorized, errors.New("request does not contain an access token"))
			context.Abort()
			return
		}
		claims, err := utils.ValidateToken(tokenString, role)
		if err != nil {
			response.Error(context, "failed", http.StatusUnauthorized, err)
			context.Abort()
			return
		}

		context.Set("claims", claims)

		context.Next()
	}
}

func CheckLoggedIn(context *gin.Context) bool {
	check := true
	authorizationToken := context.GetHeader("Authorization")
	tokenString := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))

	if strings.Contains(authorizationToken, "Bearer") == false {
		check = false
	}

	if tokenString == "" {
		check = false
	}

	return check

}
