package utils

import (
	"time"

	"github.com/AisAif/recipe-management-rest-api/src/http/resources"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func GenerateJWT(user resources.UserResource) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(30 * 24 * time.Hour)
	claims["authorized"] = true
	claims["user"] = user

	tokenString, err := token.SignedString([]byte(viper.GetString("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
