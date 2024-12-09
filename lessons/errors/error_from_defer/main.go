package main

import "database/sql"

func getBalance(database *sql.DB, clientId int) (balance float32, err error) {
	query := "..."
	var rows *sql.Rows
	rows, err = database.Query(query, clientId)
	if err != nil {
		return 0, err
	}

	defer func() {
		if err == nil {
			err = rows.Close()
		}
	}()

	// reading...
	return balance, err
}
