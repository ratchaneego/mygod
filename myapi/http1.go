package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func todosHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(w, "error : %v", err)
			return
		}
		fmt.Printf("body : %s\n", body)

		fmt.Fprintf(w, "hello %s created todos", req.Method)
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
