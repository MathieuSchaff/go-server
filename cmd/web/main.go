package main

import (
	"fmt"
	"net/http"

	"mathieu.test/pkg/handlers"
)

const portNUmber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	// fmt.Println(fmt.Sprintf("Starting ap on port %s", portNUmber))
	fmt.Printf("Starting ap on port %s", portNUmber)
	http.ListenAndServe(portNUmber, nil)
}
