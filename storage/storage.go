package storage

import (
	"log"
)

func insertEvent(active bool) {
	connect()

	stmt, err := db.Prepare("INSERT INTO events(active) VALUES(?)")
	if err != nil {
		log.Output(1, "DT #5")
		log.Fatal(err)
	}
	res, err := stmt.Exec(active)
	if err != nil {
		log.Output(1, "DT #4")
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

func getEvents() {
	var (
		id     int
		active bool
	)
	rows, err := db.Query("select id, name from users where id = ?", 1)
	if err != nil {
		log.Output(1, "DT #1")
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &active)
		if err != nil {
			log.Output(1, "DT #2")
			log.Fatal(err)
		}
		log.Println(id, active)
	}
	err = rows.Err()
	if err != nil {
		log.Output(1, "DT #3")
		log.Fatal(err)
	}

	event := &Event{
		ID:     id,
		Active: active,
	}

	log.Println(event)

	defer db.Close()
}
