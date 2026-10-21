package ttl

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestTtlSetGet_ExpiredKey(t *testing.T) {
	ttlStore := NewTtlStore(time.Microsecond * 50)
	ttlStore.Set("KeyA", "Value1")

	time.Sleep(time.Microsecond * 100)

	if val, err := ttlStore.Get("KeyA"); val != "" || !errors.Is(err, ErrNotFound) {
		t.Error("Get() should return error and empty string after TTL expired")
	}
}

func TestTtl_ExpiredKeyEvicted(t *testing.T) {
	ttlStore := NewTtlStore(time.Microsecond * 50)
	ttlStore.Set("KeyA", "Value1")

	time.Sleep(time.Microsecond * 100)

	want := []string{}
	got := ttlStore.Keys()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Keys() =%v, want %v", got, want)
	}
}

func TestTtlSetGet_BeforeExpiry(t *testing.T) {
	ttlStore := NewTtlStore(time.Microsecond * 100)
	ttlStore.Set("KeyA", "Value1")

	if val, err := ttlStore.Get("KeyA"); val == "" || err != nil {
		t.Errorf("Get() should return value, when not expired")
	}
}
