package controllers

import (
	"io"
	"log"
	"net/http"

	"github.com/carlosSimplicio/go-auth-api/src/services"
)

type AuthenticationControllerType struct {
	routeTable map[string]func(http.ResponseWriter, *http.Request)
}

func (c *AuthenticationControllerType) SetupRoutes(handler *http.ServeMux) {
	for route, routeHandler := range c.routeTable {
		handler.HandleFunc(route, routeHandler)
	}
}

func handleSignUp(w http.ResponseWriter, r *http.Request) {
	var err error
	var body []byte
	body, err = io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = services.SignUp(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var err error
	var body []byte
	body, err = io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln("Failed to read body")
	}

	token, err := services.Login(body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))
}

var AuthenticationController = &AuthenticationControllerType{
	map[string]func(http.ResponseWriter, *http.Request){
		"POST /login":  handleLogin,
		"POST /signup": handleSignUp,
	},
}
