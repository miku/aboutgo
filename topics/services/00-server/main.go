package main

import (
	"log"
	"net/http"
)

func main() {

	log.Println("http://localhost:8000")
	// Start a server listening on port 8000 and responding using Echo.
	http.ListenAndServe("localhost:8000",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello World!\n"))
		}),
	)

}
