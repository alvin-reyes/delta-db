package main

import (
	"github.com/nsqio/go-nsq"
	"log"
)

func main() {
	// Instantiate a new producer that publishes to the specified topic )
	producer, err := nsq.NewProducer("nsq-test.estuary.tech:4150", nsq.NewConfig())
	if err != nil {
		log.Fatal(err)
	}

	// Publish a message to the topic
	message := []byte("Hello, NSQ!")
	err = producer.Publish("my_topic", message)
	if err != nil {
		log.Fatal(err)
	}

	// Gracefully stop the producer when done
	producer.Stop()
}
