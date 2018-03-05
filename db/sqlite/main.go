package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	for _, s := range sql.Drivers() {
		fmt.Println(s)
	}

	db, err := sql.Open("sqlite3", "test.sqlite3.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// rows, err := db.Query(`SELECT * from WorldDirectory where name LIKE ?;`, "world")
	rows, err := db.Query(`SELECT * from WorldDirectory where name = ?`, "world1")
	if err != nil {
		log.Fatal(err)
	}

	// print the columns
	cols, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range cols {
		fmt.Print(c, "|")
	}
	fmt.Println()

	for rows.Next() {
		worldId := ""
		name := ""
		hostport := ""
		err := rows.Scan(&worldId, &name, &hostport)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(worldId, name, hostport)
	}
}
