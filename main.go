package main

import (
	"fmt"
	"net/http"
)

const portNUmber = ":8080"

// Home is the about page handler

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is my Home page")

}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(3, 4)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("this is 3 + 4 %d", sum))
}
func AddValues(x, y int) int {
	return x + y
}
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println(fmt.Sprintf("Starting ap on port %s", portNUmber))
	http.ListenAndServe(portNUmber, nil)
}
