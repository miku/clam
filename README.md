README
======

clam is a shell utility. It will run templated commands.

Examples
--------

Running a *simple* command:

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

    $ go run examples/simple.go
    2015/07/01 02:37:36 echo Hello World
    Hello World

----

Running with a *pipe* and a temporary output:

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

Running the above will create a temporary file:

    $ go run examples/pipe.go
    2015/07/01 02:41:11 echo Hello,World,! | cut -d, -f2 > /tmp/clam-370786565

    $ cat /tmp/clam-370786565
    World

----

Running with a pipe and a temporary output, this time, we want the filename returned to our program:

    package main

    import (
        "log"

        "github.com/miku/clam"
    )

    func main() {
        output, _ := clam.RunOutput("echo Hello,World,! | cut -d, -f2 > {{ output }}", clam.Map{})
        log.Printf("Find output at %s", output)
    }

Running the above will create a temporary file:

    $ go run examples/withoutput.go
    2015/07/01 02:46:55 echo Hello,World,! | cut -d, -f2 > /tmp/clam-558261601
    2015/07/01 02:46:55 Find output at /tmp/clam-558261601
