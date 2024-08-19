package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	mysql "github.com/carlosSimplicio/go-auth-api/src/mysql"
	"golang.org/x/crypto/bcrypt"
)

type SignInData struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
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

		hashedPassword, err := hashPassword([]byte(data.Password))
		if err != nil {
			log.Fatalf("Failed to hash password")
		}

		_, err = mysql.Exec(
			"INSERT INTO user (name, email, password) VALUES (?,?,?)",
			&data.Name, 
			&data.Email, 
			&hashedPassword,
		)

		if err != nil {
			log.Fatalln("Failed to insert user Id")
		}

		w.Write([]byte("User created successfully"))
	})

	mysql.Connect()
	fmt.Println("Starting server at port:", PORT)
	log.Fatal(server.ListenAndServe())
}

func hashPassword(password []byte) (hashedPassword []byte, err error){
	rounds := 10
	return bcrypt.GenerateFromPassword(password, rounds)
}
