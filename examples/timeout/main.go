package main

import (
	"log"
	"time"

	"github.com/miku/clam"
)

func main() {
	r := clam.NewRunnerTimeout(50 * time.Millisecond)
	err := r.Run("sleep 1", clam.Map{})
	if err != nil {
		log.Fatal(err)
	}
}
