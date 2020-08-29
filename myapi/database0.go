package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("database")
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db, err := sql.Open("postgres", "postgres://eihytncu:AJB4RN8B9nia4vAZwaqWAKJ3G36OyjB8@lallah.db.elephantsql.com:5432/eihytncu")

	if err != nil {
		log.Fatal("Connect to database error", err)
	}

	defer db.Close()
	log.Println("OKay")

	createTb := `
	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title TEXT,
		status TEXT
	);
`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table", err)
	}
	fmt.Println("create table success")
}
