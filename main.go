package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type Personalizations struct {
	To          []string
	DynamicData map[string]interface{}
	TemplateID  string
}

func main() {
	// Make a new reader that consumes from `test-topic`, partition 0.
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-topix",
		GroupID: "group-id",
	})
	defer r.Close()

	// Read a message
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("error while receiving message: %s", err.Error())
			continue
		}
		fmt.Printf("Message: of irfaq is %s\n", string(m.Value))
		// Process the message here
		// Optionally commit the offset
	}
}
