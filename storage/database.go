package storage

import (
	"database/sql"
	"fmt"
	"log"

	// Use postgresql database
	_ "github.com/lib/pq"
)

const (
	host     = "dumbo.db.elephantsql.com"
	port     = 5432
	user     = "mjpsosim"
	password = "zA9l-RqleaMcK34hc-cRyNnCnn4uq6tr"
	dbname   = "mjpsosim"
)

var db *sql.DB

func connect() {

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	//defer db.Close()
}

func createTables() {
	var createTableEvents = `CREATE TABLE events (
		ID          integer CONSTRAINT firstkey PRIMARY KEY,
		title       varchar(40) NOT NULL,
		active      integer NOT NULL,
		date_prod   date
		);`

	stmt, err := db.Prepare(createTableEvents)

	if err != nil {
		log.Output(1, "DT #6")
		log.Fatal(err)
	}
	res, err := stmt.Exec()
	if err != nil {
		log.Output(1, "DT #7")
		log.Fatal(err)
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)
}
