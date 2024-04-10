package main

import (
	"github.com/sqlc-dev/plugin-sdk-go/codegen"

	hello "example-sqlc-gen-plugin/internal"
)

func main() {
	codegen.Run(hello.Handler)
}
