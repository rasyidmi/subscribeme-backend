package helper

import (
	"errors"
	"projects-subscribeme-backend/config"
	"projects-subscribeme-backend/models"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWTClaim struct {
	Nama     string         `json:"nama"`
	Username string         `json:"username"`
	Npm      string         `json:"npm"`
	Jurusan  models.Jurusan `json:"jurusan"`
	Role     string         `json:"role"`
	jwt.StandardClaims
}

func GenerateJWT(ssoResponse models.ServiceResponse, role string) (tokenString string, err error) {
	configJwt := config.LoadAuthConfig()
	expirationTime := time.Now().Add(48 * time.Hour)
	claims := &JWTClaim{
		Nama:     ssoResponse.AuthenticationSuccess.Attributes.Nama,
		Username: ssoResponse.AuthenticationSuccess.User,
		Npm:      ssoResponse.AuthenticationSuccess.Attributes.Npm,
		Jurusan:  ssoResponse.AuthenticationSuccess.Attributes.Jurusan,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(configJwt.Secret))
	return
}

func ValidateToken(signedToken string, role string) (*JWTClaim, error) {
	configJwt := config.LoadAuthConfig()
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(configJwt.Secret), nil
		},
	)

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return claims, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return claims, err
	}
	if role != "" {
		if claims.Role != role {
			err = errors.New("you dont have authorization")
			return claims, err
		}
		return claims, err
	}
	return claims, err
}

func GetTokenClaims(ctx *gin.Context) *JWTClaim {
	claims := ctx.MustGet("claims").(*JWTClaim)
	return claims

}
