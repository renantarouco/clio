package kv

type Backend interface {
	Load() error
	Set(string, string) error
	Get(string) (string, error)
}
