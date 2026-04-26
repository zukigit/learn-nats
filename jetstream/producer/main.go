package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/zukigit/learn-nats/lib"
)

func main() {
	nc, err := nats.Connect(lib.NatsURL)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatalf("Error creating JetStream context: %v", err)
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"orders.*"},
		MaxAge:   24 * time.Hour,
	})
	if err != nil && err != nats.ErrStreamNameAlreadyInUse {
		log.Fatalf("Error creating stream: %v", err)
	}

	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("it is order number %d", i)
		_, err = js.Publish("orders.new", []byte(message))
		if err != nil {
			fmt.Printf("Failed to send message: %s, err: %v\n", message, err)
		} else {
			fmt.Printf("Sent message: %s\n", message)
		}
	}

	log.Println("Finished publishing all messages")
}
