package main

import (
	"fmt"
	"os"

	hello "example-sqlc-gen-plugin/internal"
)

func main() {
	if err := hello.Run(hello.Handler); err != nil {
		fmt.Fprintf(os.Stderr, "error generating output: %s", err)
		os.Exit(2)
	}
}
