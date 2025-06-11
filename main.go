package main

import (
	"net/http"
    "github.com/gorilla/mux"
)
var route = mux.NewRouter()

func main() {
	route.HandleFunc("/posts", add_posts).Methods("POST")
    route.HandleFunc("/posts/{id}", edit_posts).Methods("PUT")
    route.HandleFunc("/posts/{id}", delete_post).Methods("DELETE")
    route.HandleFunc("/posts", get_posts).Methods("GET")
    route.HandleFunc("/posts/{id}", get_post_by_id).Methods("GET")
	http.ListenAndServe(":8080", route)
}
