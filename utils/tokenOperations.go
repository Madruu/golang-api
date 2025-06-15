package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)
//Testing
var secretKey = []byte("secretPassword")

// Generating a token from the userId as part of the claims
func GenerateToken(clientId uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["client_id"] = clientId
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() //Token valid for two hours

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// Validating the token (JWT validate)
// returns the mapclaims type
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	//Parsing token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Checking sign method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing method")
		}
		return secretKey, nil
	})

	//Checking errors
	if err != nil {
		return nil, err
	}

	//validating token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("Invalid token")
}
