package main

import "fmt"

type S map[string]string

func (s S) Get(name string) string {
	return s[name]
}

type X struct {
	s S
}

func main() {
	x := &X{s: map[string]string{"a": "b", "f": "g"}}
	fmt.Println(x.s.Get("a"))
}
