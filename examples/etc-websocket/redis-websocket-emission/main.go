package main

import (
	"github.com/chuckpreslar/emission"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"fmt"
)

const CHANNEL_NAME = "i_dont_know"

func subscribe(emitter *emission.Emitter) {
	conn, _ := redis.Dial("tcp", ":6379")
	channel := redis.PubSubConn{conn}
	channel.Subscribe(CHANNEL_NAME)
	for {
		reply := channel.Receive()
		switch parsed := reply.(type) {
		case redis.Message:
			message := string(parsed.Data)
			emitter.Emit("message", message)
		}
	}
}

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func socketHandler(emitter *emission.Emitter) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		socket, err := upgrader.Upgrade(response, request, nil)
		fmt.Println("access")
		if err != nil {
			log.Println(err)
			return
		}

		channel := make(chan string)
		handler := func(message string) {
			channel <- message
		}
		emitter.AddListener("message", handler)
		for message := range channel {
			err := socket.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				break
			}
		}
		emitter.RemoveListener("message", handler)
		socket.Close()
	}
}

func main() {

	emitter := emission.NewEmitter()
	emitter.SetMaxListeners(-1)

	go subscribe(emitter)
	handler := socketHandler(emitter)

	http.Handle("/", handler)
	http.ListenAndServe(":5000", nil)

}