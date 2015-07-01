README
======

clam is a shell utility. It will run templated commands.

[![Build Status](https://travis-ci.org/miku/clam.svg?branch=master)](https://travis-ci.org/miku/clam)

![6943](http://etc.usf.edu/clipart/6900/6943/clam-shell_6943_sm.gif)

Examples
--------

A simple command
----------------

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

Output would be:

    $ go run examples/simple/main.go
    2015/07/01 02:37:36 echo Hello World
    Hello World

Command with pipe
-----------------

Running with a *pipe* and a temporary output:

    package main

    import (
        "log"

        "github.com/miku/clam"
    )

    func main() {
        clam.Run("echo Hello,World,! | cut -d, -f2 > {{ output }}", clam.Map{})
    }

Running the above will create a temporary file:

    $ go run examples/pipe/main.go
    2015/07/01 02:41:11 echo Hello,World,! | cut -d, -f2 > /tmp/clam-370786565

    $ cat /tmp/clam-370786565
    World

Catching the output
-------------------

Running with a pipe and a temporary output, this time, we want the filename returned to our program.

    package main

    import (
        "log"

        "github.com/miku/clam"
    )

    func main() {
        output, _ := clam.RunOutput("echo A,B | cut -d, -f2 > {{ output }}", clam.Map{})
        log.Printf("Find output at %s", output)
    }

Running the above will create a temporary file:

    $ go run examples/withoutput/main.go
    2015/07/01 02:46:55 echo A,B | cut -d, -f2 > /tmp/clam-558261601
    2015/07/01 02:46:55 Find output at /tmp/clam-558261601

The output can be returned as `*os.File` and `*bufio.Reader` as well with
`clam.RunFile` and `clam.RunReader`, respectively.

The `output` parameter can also be passed:

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

This will simply append to the given file:

    $ go run examples/append/main.go
    2015/07/01 09:11:22 echo Hello >> /tmp/zzzz
    2015/07/01 09:11:22 echo World >> /tmp/zzzz

    $ cat /tmp/zzzz
    Hello
    World
