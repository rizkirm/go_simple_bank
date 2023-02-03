package db

import (
	"database/sql"
	"go_simple_bank/util"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	// var err error

	// conn, err := sql.Open(dbDriver, dbSource)
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
