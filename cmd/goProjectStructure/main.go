package main

import (
	"fmt"
	"my-go-project/routes"
	"net/http"
)

func main() {
	// Setup the router with defined routes
	r := routes.SetupRoutes()

	// Start the server
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", r)
}
