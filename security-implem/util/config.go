package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	AppName          string
	Debug            bool
	Port             string
	PrivateKey       string
	PublicKey        string
	EnabelMingration bool
	DBConfig         DBConfig
	RedisConfig      RedisConfig
	Limiter          Limiter
}

type DBConfig struct {
	DBName         string
	DBUsername     string
	DBPassword     string
	DBHost         string
	DBTimeZone     string
	DBMaxIdleConns int
	DBMaxOpenConns int
	DBMaxIdleTime  int
	DBMaxLifeTime  int
}

type RedisConfig struct {
	Url      string
	Password string
	Prefix   string
}

type Limiter struct {
	RateLimit int
	Burst     int
}

func ReadConfig() (Configuration, error) {
	var config Configuration
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return Configuration{}, fmt.Errorf("error reading config file: %w", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return Configuration{}, fmt.Errorf("unable to decode config into struct: %w", err)
	}

	return config, nil
}
