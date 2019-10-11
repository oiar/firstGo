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
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	_, err = db.Exec("CREATE DATABASE AAA;")
	if err != nil {
		fmt.Print(err.Error())
	}

	stmt, err := db.Prepare("CREATE TABLE AAA.user (id int NOT NULL AUTO_INCREMENT, name TEXT(10), age varchar(40), PRIMARY KEY (id));")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Printf("User Table successfully ....")
	}

	_, err = db.Exec(`INSERT INTO AAA.user ( name, age) VALUES ("abcd", '30');`)
	if err != nil {
		fmt.Println(err.Error())
	}

	var (
		id   int
		name string
		age  int
	)

	result, err := db.Query("SELECT * FROM AAA.user;")

	for result.Next() {
		err = result.Scan(&id, &name, &age)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result, id, name, age)
	}
}
