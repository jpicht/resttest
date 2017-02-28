package main

import (
	"fmt"
	"github.com/jpicht/logger"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func fail(f string, s ...string) {
	fmt.Fprintf(os.Stderr, f+"\n", s)
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Syntax:\n    %s <testfile.yaml>", os.Args[0])
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fail("Cannot read file: %s", err.Error())
	}

	r := &recipe{Actions: make([]*action, 0)}
	err = yaml.Unmarshal(data, r)
	if err != nil {
		fail("Cannot parse YAML: %s", err.Error())
	}

	l := logger.NewStderrLogger()
	r.Run(l)
}
