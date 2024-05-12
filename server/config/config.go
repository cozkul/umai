package config

import (
	"os"
	"strconv"
	"time"
)

type ConfigInstance struct {
	DBUserName     string `mapstructure:"DB_USER"`
	DBUserPassword string `mapstructure:"DB_PASSWORD"`
	DBName         string `mapstructure:"DB_NAME"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`

	Salt string `mapstructure:"PASS_SALT"`

	JwtSecret    string        `mapstructure:"JWT_SECRET"`
	JwtExpiresIn time.Duration `mapstructure:"JWT_EXPIRY"`
	JwtMaxAge    int           `mapstructure:"JWT_MAXAGE"`
}

var Config ConfigInstance

func LoadConfig() error {
	s := os.Getenv("JWT_EXPIRY")
	jwtExpiresIn, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	s = os.Getenv("JWT_MAXAGE")
	jwtMaxAge, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return err
	}

	Config = ConfigInstance{
		DBUserName:     os.Getenv("DB_USER"),
		DBUserPassword: os.Getenv("DB_PASSWORD"),
		DBName:         os.Getenv("DB_NAME"),
		DBHost:         os.Getenv("DB_HOST"),
		DBPort:         os.Getenv("DB_PORT"),
		Salt:           os.Getenv("PASS_SALT"),
		JwtSecret:      os.Getenv("JWT_SECRET"),
		JwtExpiresIn:   time.Duration(jwtExpiresIn),
		JwtMaxAge:      int(jwtMaxAge),
	}

	return nil
}
