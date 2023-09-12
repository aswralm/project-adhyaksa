package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	App
	DBQ
}

type App struct {
	UserDB     string `env:"DB_USER"`
	PasswordDB string `env:"DB_PASS"`
	NameDB     string `env:"DB_NAME"`
	PortDB     string `env:"DB_PORT"`
	HostDB     string `env:"DB_HOST"`
	LocationDB string `env:"DB_LOCATION"`
	Port       string `env:"PORT"`
}

type DBQ struct {
	CustomTime string `env:"CUSTOM_TIMEOUT"`
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Config error : ", err)
	}

	cfg := &Config{}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
