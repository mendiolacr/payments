package main

import (
	"fmt"
	"log"
	"net/http"
	"payment_platform/process"
	"payment_platform/utils"
)

func main() {
	// Initialize the database connection
	utils.InitializeDB()
	defer utils.CloseDB()

	// Create a new router
	r := process.SetupRouter()

	// Start server
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
