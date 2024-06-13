package models

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", viper.GetString("DB_USER"), viper.GetString("DB_PASSWORD"), viper.GetString("DB_HOST"), viper.GetString("DB_PORT"), viper.GetString("DB_NAME"))

	log.Info().Msg(dsn)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	database.AutoMigrate(&User{})

	DB = database
}
