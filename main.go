package main

import (
	"github.com/AisAif/recipe-management-rest-api/src/routes"
	"github.com/spf13/viper"
)

func main() {
	router := routes.InitRouter()

	router.Run(":" + viper.GetString("port"))
}
