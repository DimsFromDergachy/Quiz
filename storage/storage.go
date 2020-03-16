package storage

import (
	"log"
)

type events struct {
	Events []event
}

func insertEvent(event event) {
	stmt, err := db.Prepare("INSERT INTO events(title, active, date_prod) VALUES($1, $2, $3)")
	if err != nil {
		log.Output(1, "DT #5")
		log.Fatal(err)
	}
	res, err := stmt.Exec(event.title, event.active, event.date)
	if err != nil {
		log.Output(1, "DT #4")
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Affected = %d\n", rowCnt)
}

func getEvents() {
	var (
		id     int
		active bool
	)
	rows, err := db.Query("select id, active from events")
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

	event := &event{
		id:     id,
		active: active,
	}

	log.Println(event)

	defer db.Close()
}
