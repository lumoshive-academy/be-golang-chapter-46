package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func main() {
	// Header dan Payload otomatis diatur oleh jwt.NewWithClaims
	claims := jwt.MapClaims{
		"sub":   "1234567890",
		"name":  "lumoshive",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	// Membuat token dengan klaim dan algoritma enkripsi
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if token != nil {
		fmt.Println("success create token")
	}

}
