package main

import (
	"bytes"
	"fmt"

	"github.com/miku/clam"
)

func main() {
	buf := new(bytes.Buffer)
	r := clam.Runner{Stdout: buf}

	_ = r.Run("echo Hello,World,! | awk -F, '{print $2}'", clam.Map{})
	fmt.Printf("%s", buf.String())
}
