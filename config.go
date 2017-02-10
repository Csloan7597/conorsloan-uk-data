package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

const defaultConfig = "config-dev.json"

// Config config data
type Config struct {
	Port      int
	ServePath string
	DataPath  string
	Env       string
}

// NewConfigUsingFile allows custom file
func NewConfigUsingFile(path string) Config {
	file, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	buf := bytes.NewBuffer(nil)

	if _, err := io.Copy(buf, file); err != nil {
		panic(err)
	}

	var config Config

	if err := json.Unmarshal(buf.Bytes(), &config); err != nil {
		panic(err)
	}

	return config
}

// NewConfig Default config (dev)
func NewConfig() Config {
	return NewConfigUsingFile(defaultConfig)
}
