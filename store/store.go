package store

import "errors"

var ErrEmptyKey = errors.New("Empty Key")
var ErrStoreFull = errors.New("Store is full")
