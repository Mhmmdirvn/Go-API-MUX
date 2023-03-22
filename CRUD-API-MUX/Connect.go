package main

import "database/sql"

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:Muhammadirvan011206@tcp(127.0.0.1:3306)/db_crud")
	if err != nil {
		return nil, err
	}

	return db, nil
}