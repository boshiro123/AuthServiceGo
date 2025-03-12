package auth

import (
	"auth-service-go/pkg/config"
	"auth-service-go/pkg/middleware"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var salt = config.MustLoad().Secrets.Salt

func GeneratePasswordHash(password string) string {
	saltedPassword := salt + password
	hash := sha256.New()
	hash.Write([]byte(saltedPassword))
	hashedPassword := hash.Sum(nil)
	return hex.EncodeToString(hashedPassword)
}

func ComparePassword(password, hashedPassword string) bool {

	saltedPassword := salt + password
	hash := sha256.New()
	hash.Write([]byte(saltedPassword))
	hashedInputPassword := hash.Sum(nil)
	return hex.EncodeToString(hashedInputPassword) == hashedPassword
}

func GenerateToken(userID uuid.UUID) (string, string, error) {
	accessTokenExpirationTime := time.Now().Add(40 * time.Minute)
	refreshTokenExpirationTime := time.Now().Add(7 * 24 * time.Hour)

	accessTokenClaims := &middleware.Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessTokenExpirationTime.Unix(),
		},
	}
	refreshTokenClaims := &middleware.Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshTokenExpirationTime.Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(config.MustLoad().Secrets.SecretKey))
	if err != nil {
		return "", "", err
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(config.MustLoad().Secrets.SecretKey))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}
