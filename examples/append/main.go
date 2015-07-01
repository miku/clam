package main

import (
	"github.com/miku/clam"
)

func main() {
	fn := "/tmp/zzzz"
	clam.Run("echo Hello >> {{ output }}", clam.Map{"output": fn})
	clam.Run("echo World >> {{ output }}", clam.Map{"output": fn})
	// $ cat /tmp/zzzz
	// Hello
	// World
}
