package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func newStore(username, password, host, dbname string) (*sql.DB, error) {
	connectString := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, host, dbname)
	return sql.Open("mysql", connectString)
}

func work(tx *sql.Tx) error {
	_, err := tx.Exec("INSERT INTO Persons (id) VALUES (1)")
	if err != nil {
		return err
	}
	_, err = tx.Exec("SELECT SLEEP(20)")
	return err
}

func main() {
	db, err := newStore("root", "apple@125", "localhost", "test")
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	toCommit := true

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	if err = work(tx); err != nil {
		log.Printf("ERROR: %v\n", err)
		toCommit = false
	}

	defer func() {
		if toCommit {
			if err := tx.Commit(); err != nil {
				log.Fatalf("committing transaction: %v", err)
			}
		} else {
			if err := tx.Rollback(); err != nil {
				log.Printf("WARNING: error rolling back transaction: %+v", err)
			}
		}
		if err := db.Close(); err != nil {
			log.Printf("WARNING: error closing DB connection: %+v", err)
		}
	}()

}
