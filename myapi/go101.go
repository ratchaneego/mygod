package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
    db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatal("Connect to database error", err)
    }
    defer db.Close()
    row := db.QueryRow("INSERT INTO todos (title, status) values ($1, $2)  RETURNING id", "buy bmw", "active")
    var id int
    err = row.Scan(&id)
    if err != nil {
        fmt.Println("can't scan id", err)
        return
    }
    fmt.Println("insert todo success id : ", id)
}