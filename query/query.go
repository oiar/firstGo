package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/?parseTime=true")
	if err != nil {
		fmt.Print(err.Error())
	}
	// defer db.Close()

	var count int
	err = db.QueryRow("select count(*) from developer_survey.developer_survey_2019;").Scan(&count)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Println(count)
	}
}
