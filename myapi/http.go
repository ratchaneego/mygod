package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("hello")
	log.Println("method:", req.Method)
	if req.Method == "GET" {
		resp := []byte(`{"name"; "kbtg get all resource"}`)
		w.Header().Add("Content-Type", "application/json")
		w.Write(resp)
		return
	}

	resp := []byte(`{"name"; "Ratchanee"}`)

	w.Header().Add("Content-Type", "application/json")
	w.Write(resp)
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Starting...")

	err := http.ListenAndServe(":1234", nil)
	//log.Println(err)
	log.Fatal(err)
	fmt.Println("End...")
}
