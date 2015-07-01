package clam

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	var cases = []struct {
		tmpl string
		ctx  Map
		err  error
	}{
		{
			tmpl: "echo Hello {{name}}",
			ctx:  Map{"name": "World"},
			err:  nil,
		},
	}

	for _, c := range cases {
		err := Run(c.tmpl, c.ctx)
		if c.err != err {
			t.Errorf("failed with %s: %+v", err, c)
		}
	}
}

func TestRunOutput(t *testing.T) {
	var cases = []struct {
		tmpl   string
		ctx    Map
		err    error
		output string
	}{
		{
			tmpl:   "echo Hello {{name}} > {{output}}",
			ctx:    Map{"name": "World"},
			err:    nil,
			output: "Hello World\n",
		},
		{
			tmpl:   "echo Hello,World | cut -d, -f2 > {{ output }}",
			ctx:    Map{},
			output: "World\n",
		},
	}

	for _, c := range cases {
		output, err := RunOutput(c.tmpl, c.ctx)
		if c.err != err {
			t.Errorf("failed with %s: %+v", err, c)
		}
		b, err := ioutil.ReadFile(output)
		if err != nil {
			t.Error(err)
		}
		if string(b) != c.output {
			t.Errorf("got %s, want %s", string(b), c.output)
		}
	}
}

func TestRunner(t *testing.T) {
	var cases = []struct {
		tmpl   string
		ctx    Map
		err    error
		output string
	}{
		{
			tmpl:   "echo Hello {{name}}",
			ctx:    Map{"name": "World"},
			err:    nil,
			output: "Hello World\n",
		},
		{
			tmpl:   "echo Hello,World | cut -d, -f2",
			ctx:    Map{},
			output: "World\n",
		},
	}
	for _, c := range cases {
		file, err := ioutil.TempFile("", "clam-test-")
		if err != nil {
			t.Error(err)
		}
		r := Runner{
			Stdout: file,
		}

		_, err = r.RunOutput(c.tmpl, c.ctx)
		if c.err != err {
			t.Errorf("failed with %s: %+v", err, c)
		}

		b, err := ioutil.ReadFile(file.Name())
		if err != nil {
			t.Error(err)
		}

		if string(b) != c.output {
			t.Errorf("got %s, want %s", string(b), c.output)
		}

		file.Sync()
	}
}

func TestRunnerTimeout(t *testing.T) {
	var r Runner
	var err error

	r = Runner{Stdout: os.Stdout, Stderr: os.Stderr, Timeout: 50 * time.Millisecond}
	_, err = r.RunOutput("sleep 1", Map{})
	switch err.(type) {
	default:
		t.Errorf("got %s, want %s", err, Timeout{})
	case Timeout:
		log.Println(err)
	}

	r = Runner{Timeout: 2 * time.Second}
	_, err = r.RunOutput("sleep 0.1", Map{})
	if err != nil {
		t.Errorf("got %s, want %v", err, nil)
	}
}

func TestRunFile(t *testing.T) {
	var cases = []struct {
		tmpl   string
		ctx    Map
		err    error
		output string
	}{
		{
			tmpl:   "echo Hello {{name}} > {{output}}",
			ctx:    Map{"name": "World"},
			err:    nil,
			output: "Hello World\n",
		},
		{
			tmpl:   "echo Hello,World | cut -d, -f2 > {{ output }}",
			ctx:    Map{},
			output: "World\n",
		},
	}

	for _, c := range cases {
		file, err := RunFile(c.tmpl, c.ctx)
		if c.err != err {
			t.Errorf("failed with %s: %+v", err, c)
		}
		rdr := bufio.NewReader(file)
		b, err := ioutil.ReadAll(rdr)
		if err != nil {
			t.Error(err)
		}
		if string(b) != c.output {
			t.Errorf("got %s, want %s", string(b), c.output)
		}
	}
}

func TestRunReader(t *testing.T) {
	var cases = []struct {
		tmpl   string
		ctx    Map
		err    error
		output string
	}{
		{
			tmpl:   "echo Hello {{name}} > {{output}}",
			ctx:    Map{"name": "World"},
			err:    nil,
			output: "Hello World\n",
		},
		{
			tmpl:   "echo Hello,World | cut -d, -f2 > {{ output }}",
			ctx:    Map{},
			output: "World\n",
		},
	}

	for _, c := range cases {
		rdr, err := RunReader(c.tmpl, c.ctx)
		if c.err != err {
			t.Errorf("failed with %s: %+v", err, c)
		}
		b, err := ioutil.ReadAll(rdr)
		if err != nil {
			t.Error(err)
		}
		if string(b) != c.output {
			t.Errorf("got %s, want %s", string(b), c.output)
		}
	}
}
