package main

import (
	"database/sql"
	"fmt"

	//"fmt"
	"log"

	_ "github.com/lib/pq"
)

func createTable() {
	//fmt.Println("database")
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

func insertTodo() {
	db, err := sql.Open("postgres", "postgres://eihytncu:AJB4RN8B9nia4vAZwaqWAKJ3G36OyjB8@lallah.db.elephantsql.com:5432/eihytncu")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}

	defer db.Close()
	row := db.QueryRow("INSERT INTO todos (title, status) values ($1, $2) RETURNING id", "buy bmw", "active")

	var id int
	err = row.Scan(&id)
	if err != nil {
		fmt.Println("can't scan id", err)
		return
	}
	fmt.Println("insert todo success id : ", id)

}

func queryOnerow() {
	db, err := sql.Open("postgres", "postgres://eihytncu:AJB4RN8B9nia4vAZwaqWAKJ3G36OyjB8@lallah.db.elephantsql.com:5432/eihytncu")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT id, title, status FROM todos where id=$1")
	if err != nil {
		log.Fatal("can'tprepare query one row statment", err)
	}
	rowId := 1
	row := stmt.QueryRow(rowId)
	var id int
	var title, status string
	err = row.Scan(&id, &title, &status)
	if err != nil {
		log.Fatal("can't Scan row into variables", err)
	}
	fmt.Println("one row", id, title, status)
}

func updateTable() {
	db, err := sql.Open("postgres", "postgres://eihytncu:AJB4RN8B9nia4vAZwaqWAKJ3G36OyjB8@lallah.db.elephantsql.com:5432/eihytncu")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()
	stmt, err := db.Prepare("UPDATE todos SET status=$2 WHERE id=$1;")

	if err != nil {
		log.Fatal("can't prepare statment update", err)
	}
	if _, err := stmt.Exec(1, "inactive"); err != nil {
		log.Fatal("error execute update ", err)
	}
	fmt.Println("update success")
}

func queryAll() {
	db, err := sql.Open("postgres", "postgres://eihytncu:AJB4RN8B9nia4vAZwaqWAKJ3G36OyjB8@lallah.db.elephantsql.com:5432/eihytncu")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, title, status FROM todos")
	if err != nil {
		log.Fatal("can't prepare query all todos statment", err)
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("can't query all todos", err)
	}
	for rows.Next() {
		var id int
		var title, status string
		err := rows.Scan(&id, &title, &status)
		if err != nil {
			log.Fatal("can't Scan row into variable", err)
		}
		fmt.Println(id, title, status)
	}
	fmt.Println("query all todos success")
}

func main() {
	//createTable()
	//insertTodo()
	//queryOnerow()
	//updateTable()
	queryAll()
}
