package auth

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(s interface{}) (string, error) {

	// Create the JWT claims
	jwtTime := os.Getenv("JWT_EXPIRATION")
	if jwtTime == "" {
		jwtTime = "24"
	}
	expiration, err := strconv.Atoi(jwtTime)
	if err != nil {
		expiration = 24
	}
	claims := jwt.MapClaims{
		"user": s,
		"exp":  time.Now().Add(time.Hour * time.Duration(expiration)).Unix(), // Token expiration time
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		secretKey = "your-secret"
	}
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT token: %v", err)
	}

	return tokenString, nil
}

// VerifyToken verifies the JWT token
func VerifyToken(tokenString string) (bool, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key
		return []byte("your-secret-key"), nil
	})

	if err != nil {
		return false, fmt.Errorf("failed to parse token: %v", err)
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Optionally, you can perform additional validation on claims here
		username := claims["username"].(string)
		fmt.Printf("Token is valid for user: %s\n", username)
		return true, nil
	}

	return false, nil
}
