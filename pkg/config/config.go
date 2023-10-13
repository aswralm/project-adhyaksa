package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Config struct {
	Db     *sql.DB
	GormDB *gorm.DB

	App
	DBQ
	Cloudinary
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

type Cloudinary struct {
	CloudName string `env:"CLOUDINARY_CLOUD_NAME"`
	ApiKey    string `env:"CLOUDINARY_API_KEY"`
	ApiScret  string `env:"CLOUDINARY_API_SECRET"`
	Folder    string `env:"CLOUDINARY_UPLOAD_FOLDER"`
}
type DBQ struct {
	CustomTime string `env:"CUSTOM_TIMEOUT"`
}

func NewConfig(path string) (*Config, error) {
	if err := godotenv.Load(fmt.Sprintf("%v/.env", path)); err != nil {
		log.Fatal("Config error : ", err)
	}

	cfg := &Config{}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
