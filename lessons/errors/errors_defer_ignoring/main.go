package main

import "database/sql"

// Need to show solution

func getBalance(database *sql.DB, clientId int) (float32, error) {
	query := "..."
	rows, err := database.Query(query, clientId)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	// reading...
	return 0., nil
}
