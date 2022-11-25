package utils

import (
	"errors"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JWTClaim struct {
	UserId string `json:"userId"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}
type RefreshClaim struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userId uuid.UUID, role string) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &JWTClaim{
		UserId: userId.String(),
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	modifiedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return modifiedToken, err
}

func GenerateRefreshToken(userId uuid.UUID) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	expirationTime := time.Now().Add(24 * 3 * time.Hour)
	claims := &RefreshClaim{
		UserId: userId.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	modifiedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return modifiedToken, nil
}

func VerifyRefreshToken(token string) (*RefreshClaim, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &RefreshClaim{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token is invalid")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := parsedToken.Claims.(*RefreshClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}
	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return nil, err
	}
	return claims, nil

}

func VerifyAccessToken(token string) (*JWTClaim, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token is invalid")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := parsedToken.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}
	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return nil, err
	}
	return claims, nil
}
