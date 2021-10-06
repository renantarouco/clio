package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

type KV struct {
	data         map[string]string
	dumpFilename string
}

func NewKV(dumpFilename string) *KV {
	return &KV{
		data: map[string]string{
			"default": "default",
		},
		dumpFilename: dumpFilename,
	}
}

func (keyvalue *KV) Load() error {
	file, err := os.OpenFile(
		keyvalue.dumpFilename,
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

		keyvalue.data[key] = value
	}

	return nil
}

func (keyvalue *KV) Set(key, value string) {
	keyvalue.data[key] = value
}

func (keyvalue KV) Get(key string) string {
	value, ok := keyvalue.data[key]

	if !ok {
		return ""
	}

	return value
}

func (keyvalue KV) Dump() error {
	kvFile, err := os.OpenFile(
		keyvalue.dumpFilename,
		os.O_CREATE|os.O_RDWR,
		os.ModePerm,
	)
	if err != nil {
		return err
	}

	defer kvFile.Close()

	for k, v := range keyvalue.data {
		kvFile.WriteString(k + "\t" + v + "\n")
	}

	return nil
}
