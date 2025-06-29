package services

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"godp.abdanhafidz.com/config"
	"godp.abdanhafidz.com/models"
	"golang.org/x/crypto/bcrypt"
)

var salt = config.Salt
var secretKey = []byte(salt)

func GenerateToken(user *models.Account) (string, error) {
	claims := models.CustomClaims{
		UserID: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Token berlaku 24 jam
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "apqobiltu.id",
		},
	}

	// Buat token dengan metode signing
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ExtractBearerToken(authHeader string) (string, error) {
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid authorization header format")
	}
	return parts[1], nil
}

func VerifyToken(bearerToken string) (uuid.UUID, string, error) {
	// fmt.Println("bearerToken :", bearerToken)

	tokenData, err := ExtractBearerToken(bearerToken)
	if err != nil {
		return uuid.UUID{}, "invalid-token", err
	} else {
		// fmt.Println("Extracted Token:", tokenData)
	}

	token, err := jwt.ParseWithClaims(tokenData, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return uuid.UUID{}, "invalid-token", err
	}

	// Extract the claims
	claims, ok := token.Claims.(*models.CustomClaims)
	if !ok || !token.Valid {
		return uuid.UUID{}, "invalid-token", err
	}
	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		return uuid.UUID{}, "expired", err
	}

	return claims.UserID, "valid", err
}

func VerifyPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("invalid password")
	}
	return nil
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
