package main

// Making a request.

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	u := "https://heise.de"
	resp, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Printf("%s %v\n", u, resp.StatusCode)
	fmt.Printf("%s %v\n", u, resp.Status)
}
