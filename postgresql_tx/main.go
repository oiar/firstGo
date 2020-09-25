package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	dbname = "transaction"
	password = "123456"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s " +
		"dbname=%s sslmode=disable password=%s",
		host, port, user, dbname, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err, "errors in OpenDB")
	}

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println(err, "errors in BeginTx")
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO tx.pets (name, species) VALUES ('Fido', 'dog'), ('Albert', 'cat')")
	if err != nil {
		tx.Rollback()
		fmt.Println(err, "errors in ExecContext for pets")
		return
	}

	// Run a query to get a count of all cats
	row := tx.QueryRow("SELECT count(*) FROM tx.pets WHERE species='cat'")
	var catCount int
	// Store the count in the `catCount` variable
	err = row.Scan(&catCount)
	if err != nil {
		tx.Rollback()
		fmt.Println(err, "errors in QueryRow")
		return
	}

	// Now update the food table, increasing the quantity of cat food by 10x the number of cats
	_, err = tx.ExecContext(ctx, "INSERT INTO tx.food (name, quantity) VALUES ('Dog Biscuit', 3), ('Cat Food', 5)")
	if err != nil {
		tx.Rollback()
		fmt.Println(err, "errors in ExecContext for food")
		return
	}

	// Commit the change if all queries ran successfully
	err = tx.Commit()
	if err != nil {
		fmt.Println(err, "errors in Commit")
	}
}
