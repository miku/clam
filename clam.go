// Package clam implements templated shell calls. It supports interop between
// go programs and the shell. You can parameterize an arbitrary shell command
// with values from you go program. If your command writes to a file, you can
// access this file instantly in you go program.
//
// Simple example:
//
//     err := clam.Run("echo Hello {{ name }}", clam.Map{"name": "World"})
//
// More examples can be found in the examples directory:
//
// https://github.com/miku/clam/tree/master/examples
//
package clam

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/fatih/color"
	"github.com/hoisie/mustache"
)

const (
	// DefaultShell is the shell program used for all commands.
	DefaultShell = "/bin/bash"
	// Version of library.
	Version = "0.1.0"
)

// Map is a shortcut for a map with string keys and values.
type Map map[string]string

// Timeout class of errors.
type Timeout struct {
	Message string
}

func (t Timeout) Error() string {
	return t.Message
}

// Runner implements various functions above cmd.Run().
type Runner struct {
	Stderr  io.Writer
	Stdout  io.Writer
	Timeout time.Duration
}

var defaultRunner = Runner{Stderr: os.Stderr, Stdout: os.Stdout}

// NewRunnerTimeout creates a new standard runner, but with a timeout.
func NewRunnerTimeout(t time.Duration) Runner {
	return Runner{Stderr: os.Stderr, Stdout: os.Stdout, Timeout: t}
}

// RunFile runs a templated command and returns the output as *os.File.
func (r Runner) RunFile(t string, m Map) (*os.File, error) {
	filename, err := r.RunOutput(t, m)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(filename)
	if err != nil {
		return file, err
	}
	return file, err
}

// RunReader runs a templated command and returns the output as *bufio.Reader.
func (r Runner) RunReader(t string, m Map) (*bufio.Reader, error) {
	file, err := r.RunFile(t, m)
	if err != nil {
		return nil, err
	}
	return bufio.NewReader(file), nil
}

// Run runs a command.
func (r Runner) Run(t string, m Map) error {
	_, err := r.RunOutput(t, m)
	return err
}

// RunOutput runs a command and returns the output filename as string.
func (r Runner) RunOutput(t string, m Map) (string, error) {
	output, ok := m["output"]
	if !ok || output == "" {
		if output == "" {
			f, err := ioutil.TempFile("", "clam-")
			if err != nil {
				return output, err
			}
			m["output"] = f.Name()
		}
	}
	c := mustache.Render(t, m)

	color.Set(color.FgGreen)
	log.Println(c)
	color.Unset()

	cmd := exec.Command(DefaultShell, "-c", c)
	cmd.Stdout = r.Stdout
	cmd.Stderr = r.Stderr

	done := make(chan bool)
	errc := make(chan error)

	if r.Timeout == 0 {
		return m["output"], cmd.Run()
	}

	go func() {
		err := cmd.Run()
		done <- true
		errc <- err
	}()
	select {
	case <-time.After(r.Timeout):
		_ = cmd.Process.Kill()
		return "", Timeout{fmt.Sprintf("timed out: %s", c)}
	case <-done:
		return m["output"], <-errc
	}
}

// Run a templated command with a given parameter map.
func Run(t string, m Map) error {
	return defaultRunner.Run(t, m)
}

// RunFile a templated command with a given parameter map. Return the output
// as file object.
func RunFile(t string, m Map) (*os.File, error) {
	return defaultRunner.RunFile(t, m)
}

// RunReader a templated command with a given parameter map. Return the output
// as a buffered reader.
func RunReader(t string, m Map) (*bufio.Reader, error) {
	return defaultRunner.RunReader(t, m)
}

// RunOutput a templated command with a given parameter map. If the parameter map
// contains a parameter named output and it's the empty string, insert a
// temporary filename.
func RunOutput(t string, m Map) (string, error) {
	return defaultRunner.RunOutput(t, m)
}
