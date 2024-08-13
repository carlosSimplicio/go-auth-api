package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	mysql "github.com/carlosSimplicio/go-auth-api/src/mysql"
)

type SignInData struct {
	Email string `json:email`
	Password string `json:password`
}


func main() {
	PORT := 8080
	handler := http.NewServeMux()
	server := &http.Server{
		Addr: fmt.Sprintf("localhost:%v", PORT),
		Handler: handler,
	}

	handler.HandleFunc("POST /signin", func(w http.ResponseWriter, r *http.Request) {
		var err error
		var body []byte
		body, err = io.ReadAll(r.Body)


		if err != nil {
			log.Fatalf("Failed to read body")
		}

		data := SignInData{}	
		err = json.Unmarshal(body, &data)

		if err != nil {
			log.Fatalf("Failed to parse JSON")
		}

		fmt.Printf("%v\n", data)
	})

	mysql.Connect()

	fmt.Println("Starting server at port:", PORT)
	log.Fatal(server.ListenAndServe())
}
