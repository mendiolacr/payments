package main

import (
	"bank_simulator/process"
	"log"
	"net/http"
)

func main() {
	r := process.SetupRouter()

	log.Println("Starting server on :8081")
	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Fatal(err)
	}
}
