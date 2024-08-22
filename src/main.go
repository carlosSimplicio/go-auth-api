package main

import (
	"fmt"
	"log"
	"net/http"

	ctrler "github.com/carlosSimplicio/go-auth-api/src/controllers"
	"github.com/carlosSimplicio/go-auth-api/src/infra/mysql"
	interfaces "github.com/carlosSimplicio/go-auth-api/src/registry"
)

var controllers = []interfaces.Controller{ctrler.AuthenticationController}

func main() {
	PORT := 8080
	handler := http.NewServeMux()
	server := &http.Server{
		Addr:    fmt.Sprintf("localhost:%v", PORT),
		Handler: handler,
	}

	mysql.Connect()
	defer mysql.Close()

	for _, controller := range controllers {
		controller.SetupRoutes(handler)
	}

	fmt.Println("Starting server at port:", PORT)
	log.Fatal(server.ListenAndServe())
}
