package main

import (
	"github.com/AisAif/recipe-management-rest-api/src/config"
	"github.com/AisAif/recipe-management-rest-api/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatal().Msg(err.Error())
	}

	utils.InitLog()

	r := gin.Default()

	r.Run(":" + viper.GetString("port"))
}
