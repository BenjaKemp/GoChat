package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRoutes registers routes to the mux.Router
func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/about", AboutHandler)
	r.HandleFunc("/product/{id}", ProductHandler) // Dynamic route with a parameter

	return r
}

// Handlers
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // Get the 'name' query parameter from the URL
	if name == "" {
		name = "Guest"
	}

	// Write the response body
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %s! Welcome to the Home Page.", name)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "This is the About Page.")
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)     // Retrieve the route parameter
	productID := vars["id"] // Get the 'id' value from the route

	// Write the response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Product ID: %s\n", productID)
}
