package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	resp, err := http.Get("http://heise.de")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := httputil.DumpResponse(resp, false) // body
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
