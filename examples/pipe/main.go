package main

import (
	"log"

	"github.com/miku/clam"
)

func main() {
	err := clam.Run("echo Hello,World,! | cut -d, -f2 > {{ output }}", clam.Map{})
	if err != nil {
		log.Fatal(err)
	}
}
