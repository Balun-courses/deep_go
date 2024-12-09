package main

import "database/sql"

type Database interface {
	Query(string) (string, error)
}

func RunQueyry(db Database, query string) {
	_, err := db.Query(query)
	if err == sql.ErrNoRows {
		// not found
	} else if err != nil {
		// error from database
	} else {
		// ok
	}
}
