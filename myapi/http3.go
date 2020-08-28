package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var todos []Todo

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func todosHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		b, err := json.Marshal(todos)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error: ", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}

	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(w, "error : %v", err)
			return
		}
		t := Todo{}
		err = json.Unmarshal(body, &t)
		if err != nil {
			fmt.Fprintf(w, "error: ", err)
		}
		todos = append(todos, t)
		fmt.Printf("% #v\n", todos)
		fmt.Fprintf(w, "hello %s created todos", "POST")
		return
	}

	resp, err := json.Marshal(todos)
	if err != nil {
		fmt.Fprintf(w, "error: ", err)
	}
	w.Header().Set("Content-Typ", "application/json")
	w.Write(resp)
	//fmt.Fprintf(w, "hello %s todos", req.Method)
}

func main() {
	http.HandleFunc("/todos", todosHandler)

	fmt.Println("Starting...")

	err := http.ListenAndServe(":1234", nil)
	//log.Println(err)
	log.Fatal(err)
	fmt.Println("End...")
}
