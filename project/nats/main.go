package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"

	nats "github.com/nats-io/nats.go"
)

func broadcastEvent(m *nats.Msg) {
	postBody, _ := json.Marshal(map[string]string{
		"user":    "bot",
		"message": string(m.Data),
	})

	bodyBuffer := bytes.NewBuffer(postBody)

	http.Post(os.Getenv("BROADCAST_URL"), "application/json", bodyBuffer)
}

func main() {
	natsUrl := os.Getenv("NATS_URL")

	nc, err := nats.Connect(natsUrl)
	if err != nil {
		log.Println(err)
	}

	_, err = nc.QueueSubscribe("todo", "todo", broadcastEvent)

	if err != nil {
		log.Println(err)
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	wg.Wait()
}
