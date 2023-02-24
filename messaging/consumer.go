package messaging

import (
	"github.com/nsqio/go-nsq"
	"log"
)

type DeltaMetricsMessageConsumer struct {
	Consumer *nsq.Consumer
}

func NewDeltaMetricsMessageConsumer() *DeltaMetricsMessageConsumer {
	consumer, err := nsq.NewConsumer(PrimaryTopic, MetricsTopicUrl, nsq.NewConfig())
	if err != nil {
		log.Fatalf("Could not create consumer: %v", err)
	}
	return &DeltaMetricsMessageConsumer{
		Consumer: consumer,
	}
}
