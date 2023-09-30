package integrationtest_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"project-adhyaksa/services/event/internal/repository/model"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rocketlaunchr/dbq"
)

const (
	queryTruncateEvent string = "ALTER TABLE events DROP FOREIGN KEY FK_events_branch_id; TRUNCATE TABLE events; TRUNCATE TABLE branchs;"
)

var (
	DB *sql.DB
)

func TestMain(m *testing.M) {
	err := godotenv.Load(os.ExpandEnv(".env.test"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	os.Exit(m.Run())
}

func database() {
	UserDB := os.Getenv("DB_USER_TEST")
	PasswordDB := os.Getenv("DB_PASS_TEST")
	NameDB := os.Getenv("DB_NAME_TEST")
	PortDB := os.Getenv("DB_PORT_TEST")
	HostDB := os.Getenv("DB_HOST_TEST")

	UserDB = url.QueryEscape(UserDB)
	PasswordDB = url.QueryEscape(PasswordDB)
	NameDB = url.QueryEscape(NameDB)
	HostDB = url.QueryEscape(HostDB)

	// Build the DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		UserDB,
		PasswordDB,
		HostDB,
		PortDB,
		NameDB,
	)

	dbConn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	dbConn.SetConnMaxLifetime(time.Minute * 1)
	dbConn.SetMaxOpenConns(10)
	dbConn.SetMaxIdleConns(10)

	DB = dbConn

	log.Println("Database connected successfully")
}

func refreshEventTable() {

	stmt, err := DB.Prepare(fmt.Sprintf(queryTruncateEvent))
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatalf("Error truncating messages table: %s", err)
	}

}

func insertBranchTable(branch model.Branch, ctx context.Context) {

	stmt := dbq.INSERT(branch.GetTableName(),
		[]string{"name", "id", "address", "created_at"},
		1,
	)

	args := []interface{}{
		branch.Name,
		branch.ID,
		branch.Address,
		branch.CreatedAt,
	}

	_, err := dbq.E(ctx, DB, stmt, nil, args)
	if err != nil {
		log.Fatalf("Error create: %s", err)
	}
	log.Println("insert data branch success")
}
