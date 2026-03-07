package main

import (
	"fmt"
	"log"
	"time"

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

	// Publish messages
	for i := 1; i <= 5; i++ {
		msg := []byte(fmt.Sprintf("Message #%d", i))
		subject := "test.subject"

		log.Printf("Publishing message %d to subject '%s': %s", i, subject, string(msg))

		err = nc.PublishMsg(&nats.Msg{
			Subject: subject,
			Data:    msg,
		})

		if err != nil {
			log.Printf("Error publishing message %d: %v", i, err)
		} else {
			log.Printf("Successfully published message %d", i)
		}

		// Wait a bit before sending the next message
		time.Sleep(1 * time.Second)
	}

	log.Println("Finished publishing all messages")
}
