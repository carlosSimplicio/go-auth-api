package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {
	PORT := 8080
	http.HandleFunc("/hello", handleHello)

	fmt.Println("Starting server at port:", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil))
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World - %v", r.Method)	
}