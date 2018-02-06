package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mgood/go-posix"
)

var (
	filename = flag.String("f", "", "Use template file instead of STDIN")
	version  = flag.Bool("version", false, "Prints the current version")
)

func Print(template []byte) ([]byte, error) {
	output, err := posix.ExpandEnv(string(template))
	if err != nil {
		return []byte{}, err
	}

	return []byte(output), nil
}

func main() {
	flag.Parse()

	if *version {
		PrintVersion(os.Stdout)
		os.Exit(0)
	}

	var err error
	var template []byte

	if *filename != "" {
		template, err = ioutil.ReadFile(*filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else {
		template, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	document, err := Print(template)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Stdout.Write(document)
}
