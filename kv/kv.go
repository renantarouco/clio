package kv

import (
	"log"
)

type KV struct {
	backend Backend
}

func NewKV(backend Backend) (*KV, error) {
	err := backend.Load()
	if err != nil {
		return nil, err
	}

	return &KV{
		backend: backend,
	}, nil
}

func (kv KV) Set(key, value string) {
	err := kv.backend.Set(key, value)
	if err != nil {
		log.Println("ERROR: %s", err)
	}
}

func (kv KV) Get(key string) string {
	value, err := kv.backend.Get(key)
	if err != nil {
		log.Println("ERROR: %s", err)

		return ""
	}

	return value
}
