package util

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/haji-sudo/ShabehRoshan/models"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateToken(user models.User) (string, error) {
	exp, err := strconv.Atoi(os.Getenv("SESSION_EXPIRE_MINUTE_TIME"))
	if err != nil {
		panic("ENV FILE EXPIRE TIME")
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID.String()
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(exp)).Unix()

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateRefreshToken(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	expDate := time.Now().Add(time.Hour * 24 * 15)
	claims["sub"] = user.ID.String()
	claims["exp"] = expDate.Unix()

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateValidationToken(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	expDate := time.Now().Add(time.Hour * 24)
	claims["sub"] = user.ID.String()
	claims["exp"] = expDate.Unix()

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		if err.Error() == "token has invalid claims: token is expired" {
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				if userID, ok := claims["sub"].(string); ok {
					return userID, fmt.Errorf("expired")
				}
			}
		}
		return "", fmt.Errorf("invalid token")
	}

	//Get the USERID from claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}

	userid, ok := claims["sub"].(string)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}

	return userid, nil
}
