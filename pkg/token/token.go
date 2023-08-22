package token

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type CustomClaims struct {
	Data string `json:"data"`
	jwt.StandardClaims
}

func GenerateToken(secretKey string, data interface{}, expireTime time.Duration) (string, error) {
	byteData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": string(byteData),
		"exp":  time.Now().Add(expireTime).Unix(),
	})

	token, err := tokenClaims.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ParseToken(tokenString string, secretKey string) (interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.Data, nil
	}

	return "", fmt.Errorf("invalid token")
}
