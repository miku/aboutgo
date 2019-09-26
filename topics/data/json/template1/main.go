package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Exercise: Load and deserialize JSON from file.

type IPInfo struct {
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Country  string `json:"country"`
	IP       string `json:"ip"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Readme   string `json:"readme"`
	Region   string `json:"region"`
	Timezone string `json:"timezone"`
}

func main() {
	// Open file, read content, deserialize into struct, then display: IP, Hostname and City.
	f, err := os.Open("hello.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var ip IPInfo

	// Variant 1: ioutil.ReadFile
	// b, err := ioutil.ReadFile("hello.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Variant 2: Unmarshal.
	// b, err := ioutil.ReadAll(f)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// if err := json.Unmarshal(b, &ip); err != nil {
	// 	log.Fatal(err)
	// }

	// Variant 3: Decoder.
	dec := json.NewDecoder(f)
	if err := dec.Decode(&ip); err != nil {
		log.Fatal(err)
	}

	fmt.Println(ip.Hostname)
}
