package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

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
	f, err := os.Open("hello.json") // or ioutil.ReadFile
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	var ip IPInfo
	if err := json.Unmarshal(b, &ip); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v %v %v\n", ip.IP, ip.Hostname, ip.City)
}
