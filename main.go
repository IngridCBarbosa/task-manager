package main

import (
	"log"
	"net/http"
)

func main() {
	initRedis()
	registerRouter()

	log.Println("server running on :8080")

	http.ListenAndServe(":8080", nil)
}
