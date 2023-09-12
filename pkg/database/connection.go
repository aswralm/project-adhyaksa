package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"project-adhyaksa/pkg/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
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
	db.SetConnMaxLifetime(time.Minute * 1)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	fmt.Println("Database connected successfully")
	return db
}
