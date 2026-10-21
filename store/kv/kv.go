package kv

import (
	"fmt"
	"sort"

	"github.com/Shinjon101/go-cache/store"
)

type Store struct {
	data    map[string]string
	maxSize int
}

func NewStore(maxSize int) *Store {
	return &Store{
		data:    make(map[string]string),
		maxSize: maxSize,
	}
}

func (s *Store) Get(key string) (string, error) {
	if key == "" {
		return "", store.ErrEmptyKey
	}
	val, ok := s.data[key]

	if !ok {
		return "", fmt.Errorf("key %s does not exists", key)
	}

	return val, nil
}

func (s *Store) Set(key string, value string) error {
	if key == "" {
		return store.ErrEmptyKey
	}
	_, exists := s.data[key]
	if s.maxSize > 0 && s.Len() >= s.maxSize && !exists {
		return fmt.Errorf("Set(%s): %w", key, store.ErrStoreFull)
	}

	s.data[key] = value
	return nil
}

func (s *Store) Delete(key string) {
	delete(s.data, key)
}

func (s *Store) Keys() []string {
	keys := make([]string, 0, len(s.data))

	for key := range s.data {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

func (s *Store) Len() int {
	return len(s.data)
}
