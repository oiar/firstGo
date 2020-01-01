package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/?parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = db.Exec("CREATE DATABASE CSV;")
	if err != nil {
		fmt.Print(err.Error())
	}

	stmt, err := db.Prepare(
		"CREATE TABLE CSV.survey_result (
			Respondent int NOT NULL AUTO_INCREMENT,
		)"
	
	)
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Print(err.Error())
	}

	_, err = db.Exec(`LOAD DATA INFILE ''
		INTO TABLE CSV
	`)
}