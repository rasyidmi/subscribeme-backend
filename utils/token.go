package utils

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	UserID   string `json:"user_id"`
	Role     string `json:"role"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(userid string, username string, role string) (tokenString string, err error) {
	expirationTime := time.Now().Add(48 * time.Hour)
	claims := &JWTClaim{
		UserID:   userid,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string, role string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
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
