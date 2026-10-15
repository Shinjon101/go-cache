package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestKeys_ReturnsSortedKeys(t *testing.T) {
	store := NewStore(5)
	store.Set("charlie", "2")
	store.Set("alpha", "1")
	store.Set("bravo", "3")

	got := store.Keys()
	want := []string{"alpha", "bravo", "charlie"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Keys() =%v, want %v", got, want)
	}
}

func TestKeysEmptyStore(t *testing.T) {
	store := NewStore(5)

	got := store.Keys()

	if len(got) != 0 {
		t.Errorf("Keys() of empty store = %v, want: empty slice", got)
	}
}

func TestGetEmpty(t *testing.T) {
	store := NewStore(5)

	if _, err := store.Get(""); err == nil || !errors.Is(err, ErrEmptyKey) {
		t.Errorf("Get() should return error for passing empty string")
	}
}

func TestSetGet_RoundTrip(t *testing.T) {
	store := NewStore(5)

	store.Set("hello", "world")
	if val, err := store.Get("hello"); err != nil || val != "world" {
		t.Errorf("Get() should return = 'world', got: %v", val)
	}
	if missingVal, err := store.Get("missing"); err == nil || missingVal != "" {
		t.Errorf("Get() should return empty string & error")
	}
}
