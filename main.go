package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNUmber = ":8080"

// Home is the about page handler

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is my Home page")

}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 10.0)
	if err != nil {
		fmt.Fprintf(w, "canot divide by 0")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))
}
func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New(("cannot divie by zero"))
		return 0, err
	}
	result := x / y
	return result, nil
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
	http.HandleFunc("/about", Divide)
	fmt.Println(fmt.Sprintf("Starting ap on port %s", portNUmber))
	http.ListenAndServe(portNUmber, nil)
}
