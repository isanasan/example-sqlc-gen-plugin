package main

import (
	"bufio"
	"bytes"
	"example-sqlc-gen-plugin/codegen/plugin"
	"fmt"
	"io"
	"os"
)

func handler(request *plugin.CodeGenRequest) (*plugin.CodeGenResponse, error) {
	var files []*plugin.File
	header := bytes.NewBuffer(nil)
	header.WriteString("hello\n")
	querier := bytes.NewBuffer(nil)
	querier.WriteString("world\n")

	files = append(files, &plugin.File{Name: "hello_world.txt", Contents: append(header.Bytes(), querier.Bytes()...)})

	return &plugin.CodeGenResponse{
		Files: files,
	}, nil
}

type Handler func(*plugin.CodeGenRequest) (*plugin.CodeGenResponse, error)

func run(h Handler) error {
	var req plugin.CodeGenRequest
	reqBlob, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	if err := req.UnmarshalVT(reqBlob); err != nil {
		return err
	}
	resp, err := h(&req)
	if err != nil {
		return err
	}
	respBlob, err := resp.MarshalVT()
	if err != nil {
		return err
	}
	w := bufio.NewWriter(os.Stdout)
	if _, err := w.Write(respBlob); err != nil {
		return err
	}
	if err := w.Flush(); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := run(handler); err != nil {
		fmt.Fprintf(os.Stderr, "error generating output: %s", err)
		os.Exit(2)
	}
}
