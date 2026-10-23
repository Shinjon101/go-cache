package main

import (
	"encoding/base64"
	"fmt"

	"github.com/Shinjon101/go-cache/middlewares"
	"github.com/Shinjon101/go-cache/store"
	"github.com/Shinjon101/go-cache/store/kv"
)

func GetEncryptedValue(storer store.Storer, key string, value string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	if err := storer.Set(key, encoded); err != nil {
		return "", fmt.Errorf("Error while setting value %v", err)
	}
	return storer.Get(key)
}

func main() {

	plain := kv.NewStore(20)
	logger := middlewares.NewLoggingMiddleware(plain)

	val, err := GetEncryptedValue(logger, "KeyA", "Koyel")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)

}
