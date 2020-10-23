package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	_ "github.com/stretchr/testify/require"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	// 😱🚫😱🚫😱🚫😱🚫😱🚫😱🚫
	dbSource = "postgresql://postgres:secret@localhost:5432/postgres?sslmode=disable"
	// 	 dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
