package configs

import "github.com/spf13/viper"

var config *Config

func Init() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	config = new(Config)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	config.Service.SecretJWT = viper.GetString("SERVICE_SECRET_JWT")
	config.Database.DSN = viper.GetString("DATABASE_DSN")

	return nil
}

func Get() *Config {
	if config == nil {
		config = new(Config)
	}

	return config
}
