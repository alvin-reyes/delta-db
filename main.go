package main

import (
	"github.com/application-research/delta-db/db_models"
	"github.com/application-research/delta-db/messaging"
	"github.com/nsqio/go-nsq"
	"log"
	"time"
)

func main() {
	// Instantiate a new producer that publishes to the specified topic
	producer, err := nsq.NewProducer(messaging.MetricsTopicUrl, nsq.NewConfig())
	if err != nil {
		log.Fatal(err)
	}

	// Publish a message to the topic
	data := db_models.Content{
		ID:                1,
		Name:              "",
		Size:              0,
		Cid:               "",
		RequestingApiKey:  "",
		PieceCommitmentId: 0,
		Status:            "",
		ConnectionMode:    "",
		LastMessage:       "",
		CreatedAt:         time.Time{},
		UpdatedAt:         time.Time{},
	}

	trace := messaging.NewDeltaMetricsTracer()
	msg := messaging.DeltaMetricsBaseMessage{
		ObjectType: "content",
		Object:     data,
	}

	trace.Trace(msg)

	producer.Stop()
}
