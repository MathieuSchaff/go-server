package main

import (
	"fmt"
	"net/http"
)

const portNUmber = ":8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// fmt.Println(fmt.Sprintf("Starting ap on port %s", portNUmber))
	fmt.Printf("Starting ap on port %s", portNUmber)
	http.ListenAndServe(portNUmber, nil)
}
