package main

import (
	"fmt"
)

import (
	"github.com/chuckpreslar/emission"
)

func main() {
	emitter := emission.NewEmitter()

	hello := func(to string) {
		fmt.Printf("Hello %s!\n", to)
	}

	count := func(count int) {
		for i := 0; i < count; i++ {
			fmt.Println(i)
		}
	}

	emitter.On("hello", hello).
		On("count", count).
		Emit("hello", "world").
		Emit("count", 5)

	emitter.Emit("hello", "sekai")
}
