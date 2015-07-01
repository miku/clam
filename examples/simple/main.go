package main

import (
	"log"

	"github.com/miku/clam"
)

func main() {
	err := clam.Run("echo Hello {{ name }}", clam.Map{"name": "World"})
	if err != nil {
		log.Fatal(err)
	}
}
