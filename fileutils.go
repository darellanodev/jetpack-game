package main

import (
	"embed"
	"io"
)

// for loading fonts (.ttf) and text files (.txt)
func loadStaticResource(filesystem embed.FS , path string) ([]byte, error) {
	file, err := filesystem.Open(path)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}