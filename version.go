package main

import (
	"fmt"
	"io"
)

const Version = "0.2"

var Revision string

func PrintVersion(w io.Writer) {
	rev := Revision
	if rev == "" {
		rev = "dev"
	}

	fmt.Fprintf(w, "envprint %s.%s\n", Version, rev)
}
