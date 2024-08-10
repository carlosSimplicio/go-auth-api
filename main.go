package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/carlosSimplicio/go-auth-api/src/controllers"
)


func main() {
	PORT := 8080

	controllers.SetRoutes()

	fmt.Println("Starting server at port:", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil))
}
