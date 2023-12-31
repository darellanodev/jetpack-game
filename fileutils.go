package main

import (
	"embed"
	"io"
)

// for loading fonts (.ttf) and text files (.txt)
func loadStaticResource(filesystem embed.FS , path string) []byte {
	file, err := filesystem.Open(path)
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return data
}