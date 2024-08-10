package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/carlosSimplicio/go-auth-api/src/controllers"
	mysql "github.com/carlosSimplicio/go-auth-api/src/infra/MySql"
)


func main() {
	PORT := 8080

	controllers.SetRoutes()
	mysql.Connect()

	fmt.Println("Starting server at port:", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil))
}
