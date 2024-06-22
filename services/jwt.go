package services

import (
	"athena-pos-backend/utils"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateTokenJWT(payload interface{}) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"payload": payload,
		"exp":     time.Now().Add(time.Hour).Unix(), // 1 hour for payload is enough
	})

	token, err := claims.SignedString([]byte(os.Getenv("ATHENA_JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ReadTokenJWT(payload string) (jwt.MapClaims, error, int) {

	token, _ := jwt.Parse(utils.TrimString(payload), func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ATHENA_JWT_SECRET")), nil
	})

	if !token.Valid {
		return nil, errors.New("Token Invalid"), http.StatusBadRequest
	}

	// global value result from client
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return claims, errors.New("cannot claim the payload token"), http.StatusInternalServerError
	}

	return claims, nil, http.StatusOK
}
