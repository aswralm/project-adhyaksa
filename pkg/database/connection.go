package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"project-adhyaksa/pkg/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsnCollection = func(driver string, config *config.Config) string {
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", config.LocationDB)

	return map[string]string{
		"mysql": fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
			config.UserDB,
			config.PasswordDB,
			config.HostDB,
			config.PortDB,
			config.NameDB,
			val.Encode(),
		),
	}[driver]
}

// Connect and return *sql.DB connection
func ConnectMYSQL(driver string, config *config.Config) *sql.DB {
	dsn := dsnCollection(driver, config)
	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Fatal("Database connection error ", err)
	}

	duration, err := time.ParseDuration(config.ConnectionPool.MaxTimeConnection)
	if err != nil {
		log.Fatal("Database connection error ", err)
	}
	db.SetConnMaxLifetime(duration)
	db.SetMaxOpenConns(config.ConnectionPool.MaxOpenConnection)
	db.SetMaxIdleConns(config.ConnectionPool.MaxIdleConnection)

	fmt.Println("Database connected successfully")
	return db
}

// for GORM
func ConnectGORM(driver string, config *config.Config) *gorm.DB {
	dsn := dsnCollection(driver, config)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gorm connection error ", err)
	}

	fmt.Println("Gorm connected successfully")
	return db
}
