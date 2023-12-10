package util

import (
	"log"
	"os"
)

func ReadInput() string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
