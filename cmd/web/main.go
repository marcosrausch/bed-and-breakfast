package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

// start starts the server
func start() {
	fmt.Printf("Starting application on port %s", port)
	http.ListenAndServe(port, nil)
}

func main() {
	start()
}