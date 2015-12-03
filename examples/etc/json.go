package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	b := []byte(`{"Name":"Bob","Food":"Pickle"}`)
	var m Message
	err := json.Unmarshal(b, &m)
	fmt.Println(err, m)

	n := Message{"Alice", "Hello", 1294706395881547000}
	d, err := json.Marshal(n)
	fmt.Println(d, err)

}
