package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("database")
	db, err := sql.Open("postgres", "postgres://eihytncu:AJB4RN8B9nia4vAZwaqWAKJ3G36OyjB8@lallah.db.elephantsql.com:5432/eihytncu")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	log.Println("OKay")
}
