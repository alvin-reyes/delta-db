package main

import (
	"encoding/json"
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
		ID:                0,
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

	message, err := json.Marshal(data)
	err = producer.Publish(messaging.PrimaryTopic, message)
	if err != nil {
		log.Fatal(err)
	}

	// Gracefully stop the producer when done
	producer.Stop()
}
