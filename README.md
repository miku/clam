README
======

clam is a shell utility. It will run templated commands.

[![Build Status](https://travis-ci.org/miku/clam.svg?branch=master)](https://travis-ci.org/miku/clam)

![6943](http://etc.usf.edu/clipart/6900/6943/clam-shell_6943_sm.gif)

Examples
--------

A simple command
----------------

```go
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
```

Output would be:

    $ go run examples/simple/main.go
    2015/07/01 02:37:36 echo Hello World
    Hello World

Command with pipe
-----------------

Running with a *pipe* and a temporary output:

```go
package main

import "github.com/miku/clam"

func main() {
    clam.Run("echo A,B | cut -d, -f2 > {{ output }}", clam.Map{})
}
```

Running the above will create a temporary file:

    $ go run examples/pipe/main.go
    2015/07/01 02:41:11 echo A,B | cut -d, -f2 > /tmp/clam-370786565

    $ cat /tmp/clam-370786565
    B

Catching the output
-------------------

Running with a pipe and a temporary output, this time, we want the filename returned to our program.

```go
package main

import (
    "log"

    "github.com/miku/clam"
)

func main() {
    output, _ := clam.RunOutput("echo A,B | cut -d, -f2 > {{ output }}", clam.Map{})
    log.Printf("find output at %s", output)
}
```

Running the above will create a temporary file:

    $ go run examples/withoutput/main.go
    2015/07/01 02:46:55 echo A,B | cut -d, -f2 > /tmp/clam-558261601
    2015/07/01 02:46:55 find output at /tmp/clam-558261601

The output can be returned as `*os.File` and `*bufio.Reader` as well with
`clam.RunFile` and `clam.RunReader`, respectively.

The `output` parameter can also be passed:

```go
package main

import "github.com/miku/clam"

func main() {
    fn := "/tmp/zzzz"
    clam.Run("echo Hello >> {{ output }}", clam.Map{"output": fn})
    clam.Run("echo World >> {{ output }}", clam.Map{"output": fn})
}
```

This will simply append to the given file:

    $ go run examples/append/main.go
    2015/07/01 09:11:22 echo Hello >> /tmp/zzzz
    2015/07/01 09:11:22 echo World >> /tmp/zzzz

    $ cat /tmp/zzzz
    Hello
    World

Timeouts
--------

Define a timeout in runner:

```go
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
```

This will kill the command process, if it won't exit in time:

    $ go run examples/timeout/main.go
    2015/07/01 10:53:21 sleep 1
    2015/07/01 10:53:21 timed out: sleep 1
    exit status 1

Absorb streams
--------------

You can use a buffer to directly absorb stdout or strerr:

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

Run:

    $ go run examples/buffer/main.go
    2015/07/01 11:54:23 echo Hello,World,! | awk -F, '{print $2}'
    World
