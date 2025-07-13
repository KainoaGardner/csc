package auth

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/config"
	"github.com/KainoaGardner/csc/internal/db"
	"github.com/KainoaGardner/csc/internal/types"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateToken(jwtKey string, userID string, expireTime time.Duration) (string, error) {
	tokenClaims := types.TokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParseToken(jwtKey string, tokenString string) (*types.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &types.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Incorrect token signing method %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("Invalid Token: %w", err)
	}

	tokenClaims, ok := token.Claims.(*types.TokenClaims)
	if !ok {
		return nil, fmt.Errorf("Cannot read claims")
	}

	return tokenClaims, nil
}

func GetTokenFromRequest(r *http.Request) (string, error) {
	auth := r.Header.Get("Authorization")
	if !strings.HasPrefix(auth, "Bearer ") {
		return "", fmt.Errorf("Invalid Authorization")
	}

	return strings.TrimPrefix(auth, "Bearer "), nil
}

func GetTokenClaims(jwtKey string, r *http.Request) (*types.TokenClaims, error) {
	token, err := GetTokenFromRequest(r)
	if err != nil {
		return nil, err
	}

	tokenClaims, err := ParseToken(jwtKey, token)
	if err != nil {
		return nil, err
	}

	return tokenClaims, nil

}

func CheckExpiredToken(claims *types.TokenClaims) bool {
	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		return false
	}
	return true
}

func CheckAdminRequest(client *mongo.Client, dbConfig config.DB, jwtKey string, r *http.Request) (int, error) {
	claims, statusCode, err := CheckValidAuth(client, dbConfig, jwtKey, r)
	if err != nil {
		return statusCode, err
	}

	dbUser, err := db.FindUser(client, dbConfig, claims.UserID)
	if err != nil {
		return http.StatusBadRequest, err
	}

	if !dbUser.Admin {
		return http.StatusBadRequest, fmt.Errorf("Not admin")
	}

	return http.StatusOK, nil
}

func CheckValidAuth(client *mongo.Client, dbConfig config.DB, jwtKey string, r *http.Request) (*types.TokenClaims, int, error) {
	claims, err := GetTokenClaims(jwtKey, r)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	if CheckExpiredToken(claims) {
		return nil, http.StatusUnauthorized, fmt.Errorf("Expired token")
	}

	return claims, http.StatusOK, nil
}
