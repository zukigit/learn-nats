package main

import (
	"log"

	"github.com/nats-io/nats.go"
	"github.com/zukigit/learn-nats/lib"
)

func main() {
	// Connect to NATS
	natsURL := lib.Getenv("NATS_URL", "nats://rocky10:4222")
	log.Printf("Connecting to NATS at %s", natsURL)

	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}
	defer nc.Drain()

	log.Println("Connected to NATS")

	// Subscribe to the subject
	subject := "test.subject"
	log.Printf("Subscribing to subject '%s'", subject)

	// _, err = nc.Subscribe(subject, func(msg *nats.Msg) {
	// 	log.Printf("Received message on subject '%s': %s", msg.Subject, string(msg.Data))
	// })

	// sub, err := nc.SubscribeSync(subject)
	// if err != nil {
	// 	log.Fatalf("Error subscribing to subject '%s': %v", subject, err)
	// }

	// msg, err := sub.NextMsg(10 * time.Second)
	// if err != nil {
	// 	log.Fatalf("Error receiving message: %v", err)
	// }

	msgChannel := make(chan *nats.Msg)
	sub, err := nc.ChanSubscribe(subject, msgChannel)
	if err != nil {
		log.Fatalf("Error subscribing to subject '%s': %v", subject, err)
	}
	defer sub.Unsubscribe()

	log.Println("Listening for messages... Press Ctrl+C to exit")

	// Keep the application running to receive messages
	// In a real application, you might want a more sophisticated way to keep
	// the application alive and handle graceful shutdown.
	for msg := range msgChannel {
		log.Printf("Received message on subject '%s': %s", msg.Subject, string(msg.Data))
	}
}
