README
======

clam is a shell utility. It will run templated commands.

    $ cat examples/simple.go
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

Running with a pipe and a temporary output:

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
    2015/07/01 02:41:11 echo Hello,World,! | cut -d, -f2 > /var/folders/cj/hpk8c18n19n3x56_8bk5wb0w0000gn/T/clam-370786565

    $ cat /var/folders/cj/hpk8c18n19n3x56_8bk5wb0w0000gn/T/clam-370786565
    World
