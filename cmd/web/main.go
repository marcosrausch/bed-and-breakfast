package main

import (
	"fmt"
	"net/http"

	"github.com/marcosrausch/bed-and-breakfast/handlers"
)

const port = ":8080"

// start starts the server
func start() {
	handlers.Home()
	handlers.About()
	fmt.Printf("Starting application on port %s", port)
	http.ListenAndServe(port, nil)
}

func main() {
	start()
}