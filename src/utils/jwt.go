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
	claims["exp"] = time.Now().Add(30 * 24 * time.Hour).Unix()
	claims["authorized"] = true
	claims["user"] = user

	tokenString, err := token.SignedString([]byte(viper.GetString("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (resources.UserResource, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("SECRET_KEY")), nil
		},
	)
	if err != nil {
		return resources.UserResource{}, err
	}

	user := claims["user"].(map[string]interface{})

	return resources.UserResource{
		Username: user["username"].(string),
		Name:     user["name"].(string),
	}, nil
}
