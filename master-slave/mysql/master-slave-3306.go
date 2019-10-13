package mysql

import (
	"database/sql"
	"errors"
	"time"
)

func CreateDB(db *sql.DB, createDB string) error {
	_, err := db.Exec(createDB)
	return err
}

func CreateTable(db *sql.DB, createTable string) error {
	_, err :=db.Exec(createTable)
	return err
}

func Insert(db *sql.DB, insert string, )