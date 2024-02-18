package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// AuthService defines the interface for authentication service
type AuthService interface {
	Authenticate(username, password string) (string, error)
	VerifyToken(tokenString string) (bool, error)
}

// JWTAuthService implements AuthService using JWT authentication
type JWTAuthService struct {
	// You can add fields here if needed
}

// NewJWTAuthService creates a new instance of JWTAuthService
func NewJWTAuthService() *JWTAuthService {
	return &JWTAuthService{}
}

// Authenticate verifies the username and password and generates a JWT token
func (svc *JWTAuthService) Authenticate(username, password string) (string, error) {

	// Implement your authentication logic here
	// For simplicity, we will just check if the username and password are not empty
	if username == "" || password == "" {
		return "", errors.New("username and password are required")
	}

	// Create the JWT claims
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT token: %v", err)
	}

	return tokenString, nil
}

// VerifyToken verifies the JWT token
func (svc *JWTAuthService) VerifyToken(tokenString string) (bool, error) {
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
