package main

import (
	"log"

	"github.com/miku/clam"
)

func main() {
	output, err := clam.RunOutput("echo Hello,World,! | cut -d, -f2 > {{ output }}", clam.Map{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Find output at %s", output)
}
