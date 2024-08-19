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

type SignUpData struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}
type LoginData struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Id int
	Name string
	Email string
	Password string
}

func main() {
	PORT := 8080
	handler := http.NewServeMux()
	server := &http.Server{
		Addr: fmt.Sprintf("localhost:%v", PORT),
		Handler: handler,
	}

	handler.HandleFunc("POST /signup", func(w http.ResponseWriter, r *http.Request) {
		var err error
		var body []byte
		body, err = io.ReadAll(r.Body)


		if err != nil {
			log.Fatalln("Failed to read body")
		}

		data := SignUpData{}	
		err = json.Unmarshal(body, &data)

		if err != nil {
			log.Fatalln("Failed to parse JSON")
		}

		hashedPassword, err := hashPassword([]byte(data.Password))
		if err != nil {
			log.Fatalln("Failed to hash password")
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

	handler.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
		var err error
		var body []byte
		body, err = io.ReadAll(r.Body)


		if err != nil {
			log.Fatalln("Failed to read body")
		}

		data := LoginData{}	
		err = json.Unmarshal(body, &data)

		if err != nil {
			log.Fatalln("Failed to parse JSON")
		}

		result, err := mysql.Select[User]("SELECT Id, Name, Email, Password FROM user WHERE Email = ?;", &data.Email)
		if err != nil {
			log.Fatalln("Failed to find user")
		}

		if len(result) == 0 {
			log.Fatalln("User not found")
		}

		user := result[0]

		if err := comparePassword([]byte(user.Password), []byte(data.Password)); err != nil {
			log.Fatalln("Invalid password")
		}

		w.Write([]byte("Successfully authenticated"))
	})

	mysql.Connect()
	fmt.Println("Starting server at port:", PORT)
	log.Fatal(server.ListenAndServe())
}

func hashPassword(password []byte) (hashedPassword []byte, err error){
	rounds := 10
	return bcrypt.GenerateFromPassword(password, rounds)
}

func comparePassword(hash, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}
