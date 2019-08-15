package main

import (
	"fmt"
	"log"
	"net/http"
)

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You asked to %s %s\n", r.Method, r.URL.Path)
}

func main() {

	// Convert the Echo function to a type that implements http.Handler
	h := &handler{}

	// Start a server listening on port 8000 and responding using Echo.
	if err := http.ListenAndServe("localhost:8000", h); err != nil {
		log.Fatalf("error: listening and serving: %s", err)
	}
}
