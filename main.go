package main

import (
	"encoding/base64"
	"fmt"

	"github.com/Shinjon101/go-cache/middlewares"
	"github.com/Shinjon101/go-cache/store"
	"github.com/Shinjon101/go-cache/store/kv"
)

func CacheEncryptedValue(storer store.Storer, key string, value string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	if err := storer.Set(key, encoded); err != nil {
		return "", fmt.Errorf("Error while setting value %v", err)
	}
	return storer.Get(key)
}

func main() {

	plain := kv.NewStore(20)
	logger := middlewares.NewLoggingMiddleware(plain)
	metrics := middlewares.NewMetricsMiddleware(logger)

	CacheEncryptedValue(metrics, "KeyA", "Linkin Park")
	metrics.Set("KeyB", "Pink Floyd")
	metrics.Set("KeyC", "Eagles")

	metrics.Get("KeyB")
	metrics.Get("KeyC")

	// A missing key increments the miss counter.
	if _, err := metrics.Get("KeyD"); err != nil {
		fmt.Println(err)
	}

	// The remaining operations are also tracked by the middleware.
	fmt.Println("Keys:", metrics.Keys())
	fmt.Println("Length:", metrics.Len())
	metrics.Delete("KeyA")
	fmt.Println("Length after delete:", metrics.Len())

	metrics.Report()
}
