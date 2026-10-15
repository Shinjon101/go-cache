package main

import "fmt"

func main() {

	s := NewStore(2)

	s.Set("a", "shinjon")

	val, err := s.Get("a")
	if err != nil {
	}
	fmt.Println(val)
	s.Keys()
}
