package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http.HandlerFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("Hello World")
	// })
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("Hello World")
		n, err := fmt.Fprintf(w, "Hello Waaaorld")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
	})
	http.ListenAndServe(":8080", nil)
}
