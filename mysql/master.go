package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	master := CreateCon("192.168.0.253:3308")
	slaveOne := CreateCon("192.168.0.251:3308")
	slaveTwo := CreateCon("192.168.0.254:3308")
	const count = 5000000

	master.CreateDB()
	master.CreateTable()

	slaveOne.CreateDB()
	slaveOne.CreateTable()

	slaveTwo.CreateDB()
	slaveTwo.CreateTable()

	start := time.Now()
	for i := 0; i < count; i++ {
		master.InsertData("name", "running")
		fmt.Println(i)
	}
	fmt.Println("[all insert time]", time.Now().Sub(start).Seconds())

	//start := time.Now()
	//go func() {
	//	var code int
	//	for code != 1 {
	//		code = master.QueryDataByName(name)
	//	}
	//	log.Println("[master query]", time.Now().Sub(start).Seconds())
	//}()

	//go func() {
	//	var code int
	//	for code != 1 {
	//		code = slaveOne.QueryDataByName(name)
	//	}
	//	log.Println("[slave 3307 query]", time.Now().Sub(start).Seconds())
	//}()

	//go func() {
	//	var code int
	//	for code != 1 {
	//		code = slaveTwo.QueryDataByName(name)
	//	}
	//
	//	log.Println("[slave 3308 query]", time.Now().Sub(start).Seconds())
	//}()
}

type DB struct {
	*sql.DB
}

// CreateCon create a db conn to local with given port
func CreateCon(address string) *DB {
	db, err := sql.Open("mysql", "root:123456@tcp("+address+")/?parseTime=true")
	if err != nil {
		log.Println(err)
	}

	return &DB{
		db,
	}
}

func (db *DB) CreateDB() error {
	result, err := db.Exec("CREATE DATABASE IF NOT EXISTS masterSlaveDB;")
	if affected, _ := result.RowsAffected(); affected == 0 {
		return errors.New("[create order] : create order affected 0 rows")
	}

	if err != nil {
		log.Println(err)
	}

	return nil
}

func (db *DB) CreateTable() {
	_, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS masterSlaveDB.masterSlaveTable (" +
			"id bigint(20) unsigned NOT NULL AUTO_INCREMENT, " +
			"name varchar(50) DEFAULT NULL, " +
			"hobbies varchar(200) DEFAULT NULL, " +
			"PRIMARY KEY (id)" +
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4; ")
	if err != nil {
		log.Println(err)
	}
}

func (db *DB) InsertData(name, hobbies string) {
	start := time.Now()

	_, err := db.Exec("INSERT INTO masterSlaveDB.masterSlaveTable (name, hobbies) VALUES (?,?)", name, hobbies)
	if err != nil {
		log.Println(err)
	}

	log.Println("[Insert Spend]", time.Now().Sub(start).Seconds())
}

func (db *DB) QueryDataByName(name string) int {
	result := db.QueryRow("SELECT * FROM masterSlaveDB.masterSlaveTable WHERE name = ? LIMIT 1 LOCK IN SHARE MODE", name)
	var (
		id      int64
		name1   string
		hobbies string
	)

	if err := result.Scan(&id, &name1, &hobbies); err != nil {
		return 0
	}

	return 1
}