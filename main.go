package main

import (
	"net/http"
	"xl/controller"
)

func main() {
	http.HandleFunc("/upload", controller.UploadHandler)
	http.HandleFunc("/login", controller.LoginHandler)
	http.HandleFunc("/register", controller.RegisterHandler)
	http.HandleFunc("/product/insert", controller.InsertHandler)
	http.HandleFunc("/product/update", controller.UpdateHandler)
	http.HandleFunc("/product/showall", controller.ShowAllHandler)
	http.HandleFunc("/product/show", controller.ShowHandler)
	http.HandleFunc("/product/destory", controller.DestoryHandler)
	http.ListenAndServe("localhost:8080", nil)
}
