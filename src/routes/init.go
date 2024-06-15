package routes

import (
	"github.com/AisAif/recipe-management-rest-api/src/config"
	"github.com/AisAif/recipe-management-rest-api/src/middleware"
	"github.com/AisAif/recipe-management-rest-api/src/models"
	"github.com/AisAif/recipe-management-rest-api/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func InitRouter() *gin.Engine {
	if err := config.Init(); err != nil {
		log.Fatal().Msg(err.Error())
	}

	utils.InitLog()
	models.InitDB()
	r := gin.Default()

	r.Use(middleware.GlobalErrorMiddleware())

	Auth(r.Group("/auth/"))
	User(r.Group("/users/"))

	return r
}
