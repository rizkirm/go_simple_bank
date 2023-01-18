package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:postgres@10.10.8.109:5436/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	// conn, err := sql.Open(dbDriver, dbSource)
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
