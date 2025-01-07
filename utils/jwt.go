package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = ""

func GenerateJWT(email, username string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"email": email,
		"exp": time.Now().Add(time.Minute * 9).Unix(),
	}).SignedString([]byte(secretKey))
	
}

func VerifyJWT(token string) (jwt.MapClaims, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, err
	}
	return t.Claims.(jwt.MapClaims), nil
}