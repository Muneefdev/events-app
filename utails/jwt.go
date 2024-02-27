package utails

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID int64, email string) (string, error) {
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"email":  email,
		"exp":    time.Now().Add(time.Hour * 3).Unix(),
	})
	token, err := jwt.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyJWT(token string) (jwt.MapClaims, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected signing method.")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	isValid := jwtToken.Valid
	if !isValid {
		return nil, errors.New("Invalid token")
	}

	claims, err := ExtractClaimsFromJWT(jwtToken)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func ExtractClaimsFromJWT(token *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Invalid token")
	}

	return claims, nil
}
