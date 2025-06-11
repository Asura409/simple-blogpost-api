package main

import (
	"fmt"
	"net/http"
)

func printHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
func main() {
	http.HandleFunc("/hello", printHello)
	http.ListenAndServe(":8080", nil)
}
