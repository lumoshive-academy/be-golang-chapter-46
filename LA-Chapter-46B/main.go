package main

import (
	"crypto/rand"
	"encoding/hex"
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

	secretKey := []byte("55edd4f00b704a31ae11f3e441c90b950c12d5dc041d660b6cddc8ab54c0b0e8")

	// Proses Signing Token
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Error signing token:", err)
		return
	}

	fmt.Println("Signed JWT:", signedToken)

	// generate secreate key
	generate_secreatekey, err := generateSecretKey(32) // 32 byte = 64 characters
	if err != nil {
		fmt.Println("Error generating secret key:", err)
		return
	}
	fmt.Println("Generated Secret Key:", generate_secreatekey)

}

func generateSecretKey(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
