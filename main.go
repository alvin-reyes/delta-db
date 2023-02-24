package main

import (
	"github.com/application-research/delta-db/messaging"
)

func main() {
	producer := messaging.NewDeltaMetricsMessageProducer()
	// add a handler to consume messages
	producer.Publish([]byte("Hello World!"))

}
