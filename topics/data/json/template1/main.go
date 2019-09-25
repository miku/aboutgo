package main

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

}
