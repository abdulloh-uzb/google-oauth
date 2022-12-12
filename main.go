package main

import (
	"google-oauth/controller"
	"net/http"
)

func main() {
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/google/login", controller.GoogleLogin)
	http.HandleFunc("/callback", controller.GoogleCallback)
	http.ListenAndServe(":8080", nil)
}
