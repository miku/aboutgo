package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", "https://heise.de", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatal(err)
	}
}
