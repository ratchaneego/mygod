package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func todosHandler(w http.ResponseWriter, req *http.Request) {
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
		
		fmt.Printf("body : % #v\n", t)
		fmt.Fprintf(w, "hello %s created todos", "POST")
		return
	}

	fmt.Fprintf(w, "hello %s todos", req.Method)
}

func main() {
	http.HandleFunc("/todos", todosHandler)

	fmt.Println("Starting...")

	err := http.ListenAndServe(":1234", nil)
	//log.Println(err)
	log.Fatal(err)
	fmt.Println("End...")
}
