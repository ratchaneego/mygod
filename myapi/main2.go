package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}

func getTodoHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	t, ok := todos[id]
	if !ok {
		c.JSON(http.StatusOK, gin, H{})
		return
		c.JSON(http.SttusOK, t)
	}
}

func CreateTodoHandler(c *gin.Context) {
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

func main() {
	r := gin.Default()
	r.GET("/todos", getTodosHandler)
	r.GET("/todos,:id", getTodosHandler)
	r.POST("/todos", createTodosHandler)
	r.Run(":1234")
}
