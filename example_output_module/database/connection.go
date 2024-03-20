package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Conn() (*sql.DB, error) {
	databaseurl := "postgresql://postgres:l@localhost/mydb"
	db, err := sql.Open("postgres", databaseurl)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
