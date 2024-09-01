package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var token *jwt.Token

func main() {

	secretKey := []byte("your-256-bit-secret")
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verifikasi algoritma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		fmt.Println("Error parsing token:", err)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Memeriksa kadaluarsa token
		if exp, ok := claims["exp"].(float64); ok {
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				fmt.Println("Token has expired")
				return
			}
		}
		fmt.Println("Token is valid")
		fmt.Println("User ID:", claims["sub"])
		fmt.Println("Name:", claims["name"])
	} else {
		fmt.Println("Invalid token")
	}
}
