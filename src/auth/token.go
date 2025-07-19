package auth

import (
	"devbook-api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userID
	permissions["sub"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString(config.SecretKey) // secret
}

func ValidateToken(r *http.Request) error {
	tokenStr := extractToken(r)
	token, err := jwt.Parse(tokenStr, getVerifiedKey)

	if err != nil {
		return err
	}

	if _, isOk := token.Claims.(jwt.MapClaims); isOk && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func extractToken(r *http.Request) string {
	tokenHeader := strings.Split(r.Header.Get("Authorization"), " ")

	if len(tokenHeader) == 2 {
		return tokenHeader[1]
	}

	return ""
}

func getVerifiedKey(token *jwt.Token) (interface{}, error) {
	if _, isOk := token.Method.(*jwt.SigningMethodHMAC); !isOk {
		return nil, fmt.Errorf("method of sign not expected! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

func ExtractUserId(r *http.Request) (uint64, error) {
	tokenStr := extractToken(r)
	token, err := jwt.Parse(tokenStr, getVerifiedKey)

	if err != nil {
		return 0, err
	}

	if permissions, isOk := token.Claims.(jwt.MapClaims); isOk && token.Valid {
		sub, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["sub"]), 10, 64)

		if err != nil {
			return 0, err
		}

		return sub, nil
	}

	return 0, errors.New("invalid token")

}
