package backend

import "log"

type DummyBackend struct{}

func (be DummyBackend) Load() error {
	log.Println("dummy load")

	return nil
}

func (be DummyBackend) Set(key, value string) error {
	log.Println("dummy set")

	return nil
}

func (be DummyBackend) Get(key string) (string, error) {
	log.Println("dummy get")

	return "dummy", nil
}
