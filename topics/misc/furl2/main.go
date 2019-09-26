// A sequential link checker that is reading from a file.
package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
)

func main() {

	f, err := os.Open("small.txt")
	if err != nil {
		log.Fatal(err)
	}
	br := bufio.NewReader(f)

	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		url := strings.TrimSpace(line)

		resp, err := http.Get(url)
		if err != nil {
			log.Printf("failed %s: %s", url, err)
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode < 400 {
			log.Printf("ok[%d]: %s", resp.StatusCode, url)
		} else {
			log.Printf("failed[%d]: %s", resp.StatusCode, url)
		}
	}
}
