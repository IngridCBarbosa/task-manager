package main

import "net/http"

func registerRouter() {
	http.HandleFunc("/tasks", taskHandler)
}
