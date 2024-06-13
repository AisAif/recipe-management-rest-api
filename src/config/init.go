package config

import (
	"github.com/spf13/viper"
)

func getEnv(env string, defaultValue string) string {
	if viper.GetString(env) != "" {
		return viper.GetString(env)
	}
	return defaultValue
}
func Init() error {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		viper.SetConfigFile("../.env")
		viper.AutomaticEnv()
		if err = viper.ReadInConfig(); err != nil {
			return err
		}
	}

	viper.SetDefault("port", getEnv("APP_PORT", "4000"))
	viper.SetDefault("log.level", getEnv("LOG_LEVEL", "debug"))

	return nil
}
