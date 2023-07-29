package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to a server
	nc, err := nats.Connect("nats://localhost:4222/")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer nc.Close()

	// Simple Async Subscriber
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message on [%s]: %s\n", m.Subject, string(m.Data))
	})

	// Simple Publisher
	nc.Publish("foo", []byte("Hello World"))

	// Responding to a request message
	nc.Subscribe("help", func(m *nats.Msg) {
		fmt.Println("Help")
		m.Respond([]byte("answer is 42"))
	})

	// Replies
	nc.Subscribe("help", func(m *nats.Msg) {
		nc.Publish(m.Reply, []byte("I can help!"))
	})

	// Requests
	msg, err := nc.Request("help", []byte("help me"), 10*time.Millisecond)
	if err != nil {
		fmt.Println("Error sending request:", err)
	} else {
		fmt.Printf("Response received on [%s]: %s\n", string(msg.Subject), string(msg.Data))
	}

	// Drain connection (Preferred for responders)
	nc.Drain()
}
