package main

import (
	"errors"
	"log"
)

var (
	ErrBadRequest = errors.New("bad request")
	ErrInvalid    = errors.New("invalid")
)

func run(i int) error {
	switch {
	case i%2 == 0:
		return ErrBadRequest
	case i%5 == 0:
		return ErrInvalid
	default:
		return nil
	}
}

func main() {
	switch err := run(5); err {
	case ErrBadRequest:
		log.Println("bad request")
	case ErrInvalid:
		log.Println("invalid")
	default:
		log.Println("all ok")
	}
}
