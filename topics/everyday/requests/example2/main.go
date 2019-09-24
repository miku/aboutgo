package main

// Making a request.

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	u := "https://heise.de"
	resp, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatal(err)
	}
}
