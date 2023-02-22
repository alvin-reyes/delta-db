package messaging

import (
	"github.com/nsqio/go-nsq"
	"log"
)

type DeltaMetricsMessageProducer struct {
	Producer *nsq.Producer
}

func NewDeltaMetricsMessageProducer() *DeltaMetricsMessageProducer {
	producer, err := nsq.NewProducer(MetricsTopicUrl, nsq.NewConfig())
	if err != nil {
		log.Fatalf("Could not create producer: %v", err)
	}
	return &DeltaMetricsMessageProducer{
		Producer: producer,
	}
}

func (p *DeltaMetricsMessageProducer) Publish(message []byte) error {
	err := p.Producer.Publish(PrimaryTopic, message)
	if err != nil {
		return err
	}
	return nil
}

func (p *DeltaMetricsMessageProducer) Stop(producer *nsq.Producer) {
	producer.Stop()
}
