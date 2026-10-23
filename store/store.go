package store

import "errors"

var ErrEmptyKey = errors.New("Empty Key")
var ErrStoreFull = errors.New("Store is full")

type Reader interface {
	Get(string) (string, error)
}

type Writer interface {
	Set(string, string) error
}

type Storer interface {
	Reader
	Writer
	Keys() []string
	Delete(string)
	Len() int
}
