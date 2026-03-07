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
	defer nc.Close()

	log.Println("Connected to NATS")

	// Subscribe to the subject
	subject := "test.subject"
	log.Printf("Subscribing to subject '%s'", subject)

	_, err = nc.Subscribe(subject, func(msg *nats.Msg) {
		log.Printf("Received message on subject '%s': %s", msg.Subject, string(msg.Data))
	})

	if err != nil {
		log.Fatalf("Error subscribing to subject '%s': %v", subject, err)
	}

	log.Println("Listening for messages... Press Ctrl+C to exit")

	// Keep the application running to receive messages
	// In a real application, you might want a more sophisticated way to keep
	// the application alive and handle graceful shutdown.
	select {}

	// Note: The following lines will not be reached because select{} blocks forever.
	// This is intentional to keep the consumer running.
	// If you need to add cleanup logic, you would typically use a signal
	// handler to break out of the select loop.
	// time.Sleep(10 * time.Second) // Example: keep alive for 10 seconds
	// log.Println("Finished listening")
}
