package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	client := &http.Client{}
	// return the error, so client won't attempt redirects
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for _, v := range via {
			log.Printf("request already made: %+v", v.URL)
		}
		return fmt.Errorf("halt: attempted redirect %v", req.URL.Path)
	}
	u := "http://heise.de"

	resp, err := client.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
