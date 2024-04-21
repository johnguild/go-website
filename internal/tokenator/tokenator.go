package tokenator

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/johnguild/go-website/internal/user"
)

// TODO load this through .env or OS var
var cookieTokenSecretKey = []byte("uuid-here-for-secret-key")

// generateToken generates a JWT token with user claims and expiration
func GenerateCookieWithToken(user *user.Credentials) (*http.Cookie, error) {
	if user == nil {
		return &http.Cookie{
			Name:     "token",
			Value:    "",
			Path:     "/",
			MaxAge:   -1, // Expires in 1 hour
			HttpOnly: true,
			// Secure:   true, // Set to true if using HTTPS
			// SameSite: http.SameSiteNoneMode,
		}, nil
	}

	// Create a claims struct
	claims := jwt.MapClaims{
		"exp":    time.Now().Add(time.Minute * 60).Unix(), // Token expires in 60 minutes matching the cookie expiration
		"iat":    time.Now().Unix(),                       // Issued at time
		"userId": user.Email,                              // User data
	}

	// Use HMAC signing method with SHA-256 (replace with your preferred method)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	jwtToken, err := token.SignedString(cookieTokenSecretKey)
	if err != nil {
		return nil, err
	}

	// Create a cookie
	cookie := http.Cookie{
		Name:     "token",
		Value:    jwtToken,
		Path:     "/",
		MaxAge:   3600, // Expires in 1 hour
		HttpOnly: true,
		// Secure:   true, // Set to true if using HTTPS
		// SameSite: http.SameSiteNoneMode,
	}

	return &cookie, nil
}

// validateToken validates the JWT token, checks for expiration, and returns claims
func ValidateCookieToken(tokenString string) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Method)
		}
		return cookieTokenSecretKey, nil
	})

	if err != nil {
		// Handle parsing errors (e.g., invalid token format)
		return nil, err
	}

	// Check if token is valid (not expired)
	if !token.Valid {
		return nil, fmt.Errorf("Token is invalid")
	}

	// If valid, return the claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}
	return nil, fmt.Errorf("Invalid claims format")
}

// This assumes the token is valid or already validated
// and just return the claims attribute
func GetCookieTokenClaimsValue(tokenString string, key string) string {
	claims, _ := ValidateCookieToken(tokenString)
	// do some catching and return empty if failed?
	return claims[key].(string)
}
