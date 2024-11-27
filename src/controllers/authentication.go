package controllers

import (
	"io"
	"log"
	"net/http"

	"github.com/carlosSimplicio/go-auth-api/src/infra/mysql"
	userRepository "github.com/carlosSimplicio/go-auth-api/src/repositories/user"
	"github.com/carlosSimplicio/go-auth-api/src/services/authentication"
)

type controller struct {
	routeTable map[string]func(http.ResponseWriter, *http.Request)
}

func (c *controller) SetupRoutes(handler *http.ServeMux) {
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
		w.Write([]byte(err.Error()))
		return
	}

	signUpService := &authentication.SignUpService{
		UserRepository: &userRepository.UserRepository{
			Client: &mysql.MySqlClient{},
		},
	}

	err = signUpService.SignUp(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
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

	loginService := &authentication.LoginService{
		UserRepository: &userRepository.UserRepository{
			Client: &mysql.MySqlClient{},
		},
	}

	token, err := loginService.Login(body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(token))
}

var AuthenticationController = &controller{
	map[string]func(http.ResponseWriter, *http.Request){
		"POST /login":  handleLogin,
		"POST /signup": handleSignUp,
	},
}
