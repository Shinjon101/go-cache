package main

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/Shinjon101/go-cache/store/kv"
	"github.com/Shinjon101/go-cache/store/ttl"
)

type Storer interface {
	Get(string) (string, error)
	Set(string, string) error
	Keys() []string
	Delete(string)
}

func GetEncryptedValue(storer Storer, key string, value string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	if err := storer.Set(key, encoded); err != nil {
		return "", fmt.Errorf("Error while setting value %v", err)
	}
	return storer.Get(key)
}

func main() {

	ttlStore := ttl.NewTtlStore(time.Minute * 15)
	store := kv.NewStore(3)

	valTtl, _ := GetEncryptedValue(ttlStore, "KeyA", "Koyel")
	fmt.Println(valTtl)

	val, _ := GetEncryptedValue(store, "KeyA", "Shinjon")
	fmt.Println(val)

}
