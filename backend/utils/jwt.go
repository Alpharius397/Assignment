package utils

import (
	"errors"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type UserJson struct {
	UserId string `json:"id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

const issuer = "backend-api"
const accessSubject = "ACCESS"
const refreshSubject = "REFRESH"

// Generate new JWT token
func newToken(userID string, email string, issuer string, subject string, expiry uint, secret string) (string, error) {

	claims := UserJson{
		userID, email, jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiry) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    issuer,
			Subject:   subject,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Parse and validate a JWT token
func parseToken(jwtToken string, issuer string, subject string, secret string) (*UserJson, error) {
	keyFunc := func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	}

	token, err := jwt.ParseWithClaims(jwtToken, &UserJson{}, keyFunc, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}), jwt.WithIssuer(issuer), jwt.WithSubject(subject), jwt.WithExpirationRequired())

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserJson); ok {
		return claims, nil
	}

	return nil, errors.New("unknown Claims")
}

// Wrapper on newToken for Access Token, relies on JWT_SECRET environment variable
func GetAccessToken(userID string, email string) (string, error) {
	JWT_SECRET, ok := os.LookupEnv("JWT_SECRET")

	if !ok {
		return "", &KeyNotFound{}
	}

	return newToken(userID, email, issuer, accessSubject, 5, JWT_SECRET)
}

// Wrapper on parseToken for Access Token, relies on JWT_SECRET environment variable
func ParseAccessToken(token string) (*UserJson, error) {
	JWT_SECRET, ok := os.LookupEnv("JWT_SECRET")

	if !ok {
		return nil, &KeyNotFound{}
	}

	return parseToken(token, issuer, accessSubject, JWT_SECRET)
}

// Wrapper on newToken for Refresh Token, relies on JWT_SECRET environment variable
func GetRefreshToken(userID string, email string) (string, error) {
	JWT_SECRET, ok := os.LookupEnv("JWT_SECRET")

	if !ok {
		return "", &KeyNotFound{}
	}
	return newToken(userID, email, issuer, refreshSubject, 30, JWT_SECRET)
}

// Wrapper on parseToken for Refresh Token, relies on JWT_SECRET environment variable
func ParseRefreshToken(token string) (*UserJson, error) {
	JWT_SECRET, ok := os.LookupEnv("JWT_SECRET")

	if !ok {
		return nil, &KeyNotFound{}
	}

	return parseToken(token, issuer, refreshSubject, JWT_SECRET)
}
