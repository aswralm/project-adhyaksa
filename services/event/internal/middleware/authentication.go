package middleware

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

// Helper for get user data from other service
type Auth struct {
	UserID uint
}

// User function will parse the token string and return jwt.MapClaims and error
// jwt.MapClaims contains the informations from parsed jwt
// err will be nil if all parsing process are successfully
// Auth.UserID will be set from "seq" column for ease of use
func (auth *Auth) User(tokenString, jwtSignatureKey string) (jwt.MapClaims, error) {
	// Replace Bearer if exists
	if strings.Contains(tokenString, "Bearer ") {
		tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")
	}

	// Parse the token with jwtSignatureKey
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("signing method is invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, errors.New("signing method is invalid")
		}

		return []byte(jwtSignatureKey), nil
	})
	if err != nil {
		zap.L().Error(err.Error())
		return jwt.MapClaims{}, err
	}

	// Get claims
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		err := errors.New("failed to parse claims")
		zap.L().Error(err.Error())
		return jwt.MapClaims{}, err
	}

	auth.UserID = uint(claims["seq"].(float64))

	return claims, nil
}
