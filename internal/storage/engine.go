package storage

import (
	"errors"
)

// When a get request misses
var ErrKeyNotFound = errors.New("Key not found")

type Engine interface {
	Put(key string, value []byte) error
	Get(key string) ([]byte, error)
	Delete(key string) error
}