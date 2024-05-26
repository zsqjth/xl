package main

import (
	"net/http"
	"xl/controller"
)

func main() {
	http.HandleFunc("/login", controller.LogoutHandler)
	http.HandleFunc("/register", controller.RegisterHandler)
	http.ListenAndServe("localhost:8080", nil)
}
