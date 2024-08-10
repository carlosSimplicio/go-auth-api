package controllers

import (
	"net/http"

	"github.com/carlosSimplicio/go-auth-api/src/services"
)

func SetRoutes() {
	http.HandleFunc("/authenticate", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET": 
			services.Authenticate()
		default: 
			 http.NotFound(w, r)
		}
	})

}

