package messaging

import "encoding/json"

var (
	MetricsTopicUrl = "145.40.77.207:4150"
	PrimaryTopic    = "delta-metric-events"
	producer        *DeltaMetricsMessageProducer
)

type DeltaMetricsBaseMessage struct {
	ObjectType string      `json:"object_type"`
	Object     interface{} `json:"object"`
}

func init() {
	producer = NewDeltaMetricsMessageProducer()
}

func (p *DeltaMetricsBaseMessage) Trace(message DeltaMetricsBaseMessage) error {
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	producer.Publish(messageBytes)
	return nil
}

func (p *DeltaMetricsBaseMessage) TraceMessageString(message string) error {
	err := producer.Publish([]byte(message))
	if err != nil {
		return err
	}
	return nil
}

func (p *DeltaMetricsBaseMessage) TraceMessage(message []byte) error {
	err := producer.Publish(message)
	if err != nil {
		return err
	}
	return nil
}
