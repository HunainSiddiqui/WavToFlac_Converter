package main

import (
	"log"
	"net/http"
	"task-golang/router"
)

func main() {
	r := router.SetupRouter()

	log.Println("Server running on :3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}
