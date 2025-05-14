package config

import (
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type AppConfig struct {
	App struct {
		Name string
		Port int
		Env  string
	}
	Redis struct {
		Host     string
		Password string
		Port     string
		DataBase int
	}
	Database struct {
		Host            string
		Port            string
		User            string
		Password        string
		Name            string
		SSLMode         string
		MaxIdleConns    int
		MaxOpenConns    int
		ConnMaxLifetime time.Duration
	}
	OTPCode struct {
		ExpireTime time.Duration
		TryAttempt int
	}
	JWTConfig struct {
		AccessTokenExpireDuration  time.Duration
		RefreshTokenExpireDuration time.Duration
		Secret                     string
		RefreshSecret              string
	}
	Logger struct {
		Filepath string
		Encoding string
		Level    string
		Driver   string
	}
}

var Config *AppConfig

func LoadConfig() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("خطا در خواندن فایل .env: %v", err)
	}

	Config = &AppConfig{
		App: struct {
			Name string
			Port int
			Env  string
		}{
			Name: viper.GetString("APP_NAME"),
			Port: viper.GetInt("APP_PORT"),
			Env:  viper.GetString("APP_ENV"),
		},
		Redis: struct {
			Host     string
			Password string
			Port     string
			DataBase int
		}{
			Host:     viper.GetString("REDIS_HOST"),
			Password: viper.GetString("REDIS_PASSWORD"),
			Port:     viper.GetString("REDIS_PORT"),
			DataBase: viper.GetInt("REDIS_DATABASE"),
		},
		Database: struct {
			Host            string
			Port            string
			User            string
			Password        string
			Name            string
			SSLMode         string
			MaxIdleConns    int
			MaxOpenConns    int
			ConnMaxLifetime time.Duration
		}{
			Host:            viper.GetString("POSTGRES_HOST"),
			Port:            viper.GetString("POSTGRES_PORT"),
			User:            viper.GetString("POSTGRES_USER"),
			Password:        viper.GetString("POSTGRES_PASSWORD"),
			Name:            viper.GetString("POSTGRES_DATABASE"),
			SSLMode:         viper.GetString("POSTGRES_SSL_MODE"),
			MaxIdleConns:    viper.GetInt("POSTGRES_MAX_IDLE_CONNS"),
			MaxOpenConns:    viper.GetInt("POSTGRES_MAX_OPEN_CONNS"),
			ConnMaxLifetime: time.Duration(viper.GetInt("POSTGRES_CONN_MAX_LIFE_TIME")),
		},
		OTPCode: struct {
			ExpireTime time.Duration
			TryAttempt int
		}{
			ExpireTime: time.Duration(viper.GetInt("OTP_EXPIRE_TIME")),
			TryAttempt: viper.GetInt("OTP_TRY_ATTEMPT"),
		},
		JWTConfig: struct {
			AccessTokenExpireDuration  time.Duration
			RefreshTokenExpireDuration time.Duration
			Secret                     string
			RefreshSecret              string
		}{
			AccessTokenExpireDuration:  time.Duration(viper.GetInt("JWT_ACCESS_TOKEN_EXPIRE_DURATION")),
			RefreshTokenExpireDuration: time.Duration(viper.GetInt("JWT_REFRESH_TOKEN_EXPIRE_DURATION")),
			Secret:                     viper.GetString("JWT_SECRET"),
			RefreshSecret:              viper.GetString("JWT_REFRESH_SECRET"),
		},
		Logger: struct {
			Filepath string
			Encoding string
			Level    string
			Driver   string
		}{
			Filepath: viper.GetString("LOGGER_FILEPATH"),
			Encoding: viper.GetString("LOGGER_ENCODING"),
			Level:    viper.GetString("LOGGER_LEVEL"),
			Driver:   viper.GetString("LOGGER_DRIVER"),
		},
	}
}
