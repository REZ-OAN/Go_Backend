package db

import (
	"database/sql"
	"log"
	"os"
	"simple-bank/utils"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	path := os.Getenv("path")
	name := os.Getenv("name")
	ext := os.Getenv("ext")

	if path == "" || name == "" || ext == "" {
		log.Fatalln("All arguments should be given!!!")
		return
	}
	config, err := utils.LoadConfig(path, name, ext)
	if err != nil {
		log.Fatalf("%v", err)
	}
	testDB, err = sql.Open(config.DB_DRIVER, config.DB_SOURCE)
	if err != nil {
		log.Fatal("cannot connect to the db: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
