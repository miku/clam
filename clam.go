package clam

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/hoisie/mustache"
)

const DefaultShell = "/bin/bash"

type Map map[string]string

// Run a templated command with a given parameter map.
func Run(t string, m Map) error {
	_, err := RunOutput(t, m)
	return err
}

// RunFile a templated command with a given parameter map. Return the output
// as file object.
func RunFile(t string, m Map) (*os.File, error) {
	filename, err := RunOutput(t, m)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(filename)
	if err != nil {
		return file, err
	}
	return file, err
}

// RunReader a templated command with a given parameter map. Return the output
// as a buffered reader.
func RunReader(t string, m Map) (*bufio.Reader, error) {
	file, err := RunFile(t, m)
	if err != nil {
		return nil, err
	}
	return bufio.NewReader(file), nil
}

// RunOutput a templated command with a given parameter map. If the parameter map
// contains a parameter named output and it's the empty string, insert a
// temporary filename.
func RunOutput(t string, m Map) (string, error) {
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
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return m["output"], cmd.Run()
}
