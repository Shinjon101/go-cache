package ttl

import (
	"errors"
	"sort"
	"time"

	"github.com/Shinjon101/go-cache/store"
)

var ErrNotFound = errors.New("Key not found")

type TtlEntry struct {
	value     string
	expiresAt time.Time
}

type TtlStore struct {
	data map[string]TtlEntry
	ttl  time.Duration
}

func NewTtlStore(ttl time.Duration) *TtlStore {
	return &TtlStore{
		data: make(map[string]TtlEntry),
		ttl:  ttl,
	}
}

func (t *TtlStore) Get(key string) (string, error) {
	if key == "" {
		return "", store.ErrEmptyKey
	}
	entry, ok := t.data[key]

	if !ok || time.Now().After(entry.expiresAt) {
		delete(t.data, key)
		return "", ErrNotFound
	}

	return entry.value, nil
}

func (t *TtlStore) Set(key string, value string) error {
	if key == "" {
		return store.ErrEmptyKey
	}

	t.data[key] = TtlEntry{
		value:     value,
		expiresAt: time.Now().Add(t.ttl),
	}

	return nil
}

func (t *TtlStore) Delete(key string) {
	delete(t.data, key)
}

func (t *TtlStore) Keys() []string {
	keys := make([]string, 0, len(t.data))

	for key := range t.data {
		entry, ok := t.data[key]
		if !ok || time.Now().After(entry.expiresAt) {
			delete(t.data, key)
		} else {
			keys = append(keys, key)
		}

	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

func (t *TtlStore) Len() int {
	return len(t.data)
}
