package ttl

import (
	"errors"
	"testing"
	"time"
)

func TestTtlSetGet_ExpiredKey(t *testing.T) {
	ttlStore := NewTtlStore(time.Microsecond * 2)
	ttlStore.Set("KeyA", "Value1")

	time.Sleep(time.Microsecond * 5)

	if val, err := ttlStore.Get("KeyA"); val != "" || !errors.Is(err, ErrNotFound) {
		t.Error("Get() should return error and empty string after TTL expired")
	}

}
