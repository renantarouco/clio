package backend

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

type FSBackend struct {
	filename string
	data     map[string]string
}

func NewFSBackend(filename string) *FSBackend {
	return &FSBackend{
		filename: filename,
		data:     map[string]string{},
	}
}

func (be *FSBackend) Load() error {
	file, err := os.OpenFile(
		be.filename,
		os.O_CREATE|os.O_APPEND|os.O_RDWR,
		os.ModePerm,
	)
	if err != nil {
		return err
	}

	fileReader := bufio.NewReader(file)

	for {
		line, _, err := fileReader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		lineString := string(line)
		lineString = strings.TrimPrefix(lineString, "\n")
		entry := strings.Split(lineString, "\t")
		if len(entry) < 2 {
			return errors.New("malformed kv entry")
		}

		key, value := entry[0], entry[1]

		be.data[key] = value
	}

	return nil
}

func (be *FSBackend) Set(key, value string) error {
	be.data[key] = value

	return be.dump()
}

func (be FSBackend) Get(key string) (string, error) {
	value, ok := be.data[key]

	if !ok {
		return "", nil
	}

	return value, nil
}

func (be FSBackend) dump() error {
	kvFile, err := os.OpenFile(
		be.filename,
		os.O_CREATE|os.O_RDWR,
		os.ModePerm,
	)
	if err != nil {
		return err
	}

	defer kvFile.Close()

	for k, v := range be.data {
		kvFile.WriteString(k + "\t" + v + "\n")
	}

	return nil
}
