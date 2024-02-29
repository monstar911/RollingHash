package main

import (
	"fmt"
	"log"
	"os"

	"rollinghash/pkg/rollinghash"
)

func main() {
	original, err := os.Open("test/original.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer original.Close()

	updated, err := os.Open("test/updated.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer updated.Close()

	r := rollinghash.NewRollingHash(1024)

	delta, err := r.ComputeDelta(original, updated)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(delta))
}
