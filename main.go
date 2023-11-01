package main

import "net/http"

func Create(w http.ResponseWriter, r *http.Request) {}

func Read(w http.ResponseWriter, r *http.Request) {}

func main() {
	http.HandleFunc("/users/create", Create)
	http.HandleFunc("/users/read", Read)
	http.ListenAndServe(".8000", nil)
}
