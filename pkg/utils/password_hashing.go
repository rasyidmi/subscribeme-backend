package utils

import (
	"fmt"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	min := 4
	max := 15
	randomNumber := rand.Intn(max-min) + min
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), randomNumber)
	fmt.Println("HASHEDD")
	return string(bytes), err
}

func ComparePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
