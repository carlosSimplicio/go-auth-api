package controllers

import (
	"io"
	"net/http"

	"github.com/carlosSimplicio/go-auth-api/src/services"
)

func SetRoutes() {
	http.HandleFunc("/authenticate", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET": 
			handleAuthenticate(w, r)	
		
		default: 
			w.WriteHeader(http.StatusMethodNotAllowed)	
			w.Write([]byte("Invalid method for this route"))
		}
	})

	http.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handleSignIn(w, r)
		default: 
			w.WriteHeader(http.StatusMethodNotAllowed)	
			w.Write([]byte("Invalid method for this route"))
		}
	})

}

func handleAuthenticate(w http.ResponseWriter, _ *http.Request) {
	result := services.Authenticate()

	io.WriteString(w, result)
}

func handleSignIn(w http.ResponseWriter, _ *http.Request) {
	result := services.SignIn()

	w.Write([]byte(result))	
}

