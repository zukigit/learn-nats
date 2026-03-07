package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/zukigit/learn-nats/lib"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nc, err := nats.Connect(lib.NatsURL)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}
	defer nc.Close()

	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatalf("Error creating JetStream context: %v", err)
	}

	js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"ORDERS.*"},
	})

	var message string
	for i := 0; i < 100; i++ {
		message = fmt.Sprintf("it is order number %d", i)
		_, err = js.Publish(ctx, "ORDERS.new", []byte(message))
		if err != nil {
			fmt.Printf("Failed to Sent message: %s, err: %v\n", message, err)
		} else {
			fmt.Printf("Sent message: %s\n", message)
		}

		time.Sleep(1 * time.Second)
	}

	log.Println("Finished publishing all messages")
}
