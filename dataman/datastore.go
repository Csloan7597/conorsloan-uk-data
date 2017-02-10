package dataman

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"sync"
)

// DataStore a place to store Data line by line
type DataStore interface {
	Write([]byte) error
	WriteLines([][]byte) error
	Read() ([]byte, error)
	ReadLines() ([][]byte, error) // TODO: Make this stream instead of buffer it all
	Empty() error
}

// OSFileDataStore implementation of DataStore which saves data to a file
type OSFileDataStore struct {
	path string
	lock sync.RWMutex
}

// Writes the specified content to the data store as is
func (store *OSFileDataStore) Write(bytes []byte) error {
	store.lock.Lock()
	defer store.lock.Unlock()

	file, err := os.OpenFile(store.path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(bytes)

	return err
}

// WriteLines appends the specified content to the file with a new line
func (store *OSFileDataStore) WriteLines(bytes [][]byte) error {
	store.lock.Lock()
	defer store.lock.Unlock()

	file, err := os.OpenFile(store.path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer file.Close()

	for _, line := range bytes {
		lineAndNewline := append(line, []byte("\n")...)
		if _, err := file.Write(lineAndNewline); err != nil {
			return err
		}
	}

	return nil
}

// Read copies the whole file into a buffer for now, then returns it
func (store *OSFileDataStore) Read() ([]byte, error) {
	store.lock.RLock()
	defer store.lock.RUnlock()

	file, err := os.OpenFile(store.path, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	buf := bytes.NewBuffer(nil)

	if _, err := io.Copy(buf, file); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// ReadLines uses a scanner to read the whole file line by line
func (store *OSFileDataStore) ReadLines() ([][]byte, error) {
	store.lock.RLock()
	defer store.lock.RUnlock()

	file, err := os.OpenFile(store.path, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines [][]byte

	scanner := bufio.NewScanner(file)

	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	for scanner.Scan() {
		lines = append(lines, scanner.Bytes())
	}

	for scanner.Scan() {
		lines = append(lines, scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// Empty Truncates the file
func (store *OSFileDataStore) Empty() error {
	store.lock.Lock()
	defer store.lock.Unlock()

	file, err := os.OpenFile(store.path, os.O_TRUNC, 0600)
	if err != nil {
		return err
	}

	defer file.Close()

	return file.Truncate(0)
}
