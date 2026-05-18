package utils

import (
	"fmt"
	"time"

	"errors"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "your_secret_key"

func GenerateJwtToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return 0, errors.New("Unexpected signing method")
		}

		return []byte(secretKey), nil

	})

	if err != nil {
		fmt.Println("Error parsing token:")
		return 0, errors.New("Could not parse token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		fmt.Println("Token is invalid")
		return 0, errors.New("Token is invalid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		fmt.Println("Error parsing claims")
		return 0, errors.New("Could not parse claims")
	}

	email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	fmt.Printf("Token is valid. Email: %s, UserID: %d\n", email, userId)

	return userId, nil

}
