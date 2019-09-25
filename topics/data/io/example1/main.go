package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Question: How could function Contains be improved?

func Contains(f *os.File, term string) (bool, error) {
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return false, err
	}
	return strings.Contains(string(b), term), nil
}

func main() {
	f, err := os.Open("../README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	term := "iox"

	ok, err := Contains(f, term)
	if err != nil {
		log.Fatal(err)
	}

	if ok {
		fmt.Printf("file contains %s\n", term)
	} else {
		fmt.Println("no match")
	}
}
