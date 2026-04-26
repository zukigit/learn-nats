package main

import (
	"fmt"
	"log"

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
		log.Fatalf("Error creating JetStream: %v", err)
	}

	fmt.Println("started listening")

	sub, err := js.Subscribe("orders.*", func(msg *nats.Msg) {
		log.Printf("Received message: %s\n", string(msg.Data))
		msg.Ack()
	},
		nats.BindStream("ORDERS"),
		nats.Durable("CONSUMER2"),
		nats.ManualAck(),
	)
	if err != nil {
		log.Fatalf("Error subscribing: %v", err) // what does this say?
	}
	log.Printf("Subscribed: %v", sub.Subject)

	select {}
}
