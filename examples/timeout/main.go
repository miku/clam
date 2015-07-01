package main

import (
	"log"
	"os"
	"time"

	"github.com/miku/clam"
)

func main() {
	r := clam.Runner{Stdout: os.Stdout, Stderr: os.Stderr, Timeout: 50 * time.Millisecond}
	err := r.Run("sleep 1", clam.Map{})
	if err != nil {
		log.Fatal(err)
	}
}
