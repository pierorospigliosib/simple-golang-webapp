package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Read the environment variable called TEST_ENV_TEST
	envVar := os.Getenv("TEST_ENV_TEST")
	if envVar == "" {
		envVar = "The environment variable 'TEST_ENV_TEST' is not set."
	}

	// Write the HTTP response
	html := formatTextToHTML(envVar)
	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/", handler)

	// Start the web server on port 8080
	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
