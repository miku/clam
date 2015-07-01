// $ X=Hello go run examples/envvar/main.go
// 2015/07/01 10:24:03 echo "$X"
// Hello
package main

import (
	"github.com/miku/clam"
)

func main() {
	clam.Run(`echo "$X"`, clam.Map{})
}
