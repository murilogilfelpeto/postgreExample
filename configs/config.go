package configs

import (
	"github.com/spf13/viper"
	"log"
)

var configuration *config

type config struct {
	Api      ApiConfig
	Database DbConfig
}

type ApiConfig struct {
	Port string
}

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// This init is always executed before main.
// It belongs to GO life cycle
func init() {
	viper.SetDefault("api.port", "8080")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.user", "felpeto")
	viper.SetDefault("database.password", "admin")
	viper.SetDefault("database.database", "todo")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found")
			return err
		}
	}
	configuration = new(config)
	configuration.Api = ApiConfig{
		Port: viper.GetString("api.port"),
	}
	configuration.Database = DbConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Database: viper.GetString("database.database"),
	}

	return nil
}

func GetDB() DbConfig {
	return configuration.Database
}

func GetPort() string {
	return configuration.Api.Port
}
