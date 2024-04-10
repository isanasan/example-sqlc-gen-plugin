package hello

import (
	"bytes"
	"context"

	"github.com/sqlc-dev/plugin-sdk-go/plugin"
)

func Handler(_ context.Context, request *plugin.GenerateRequest) (*plugin.GenerateResponse, error) {
	var files []*plugin.File
	header := bytes.NewBuffer(nil)
	header.WriteString("hello\n")
	querier := bytes.NewBuffer(nil)
	querier.WriteString("world\n")

	files = append(files, &plugin.File{Name: "hello_world.txt", Contents: append(header.Bytes(), querier.Bytes()...)})

	return &plugin.GenerateResponse{
		Files: files,
	}, nil
}
