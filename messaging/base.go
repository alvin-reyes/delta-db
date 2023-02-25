package messaging

import "encoding/json"

var (
	MetricsTopicUrl = "145.40.77.207:4150"
	PrimaryTopic    = "delta-metric-events"
	producer        *DeltaMetricsMessageProducer
)

type DeltaMetricsTracer struct {
	Producer *DeltaMetricsMessageProducer
}
type DeltaMetricsBaseMessage struct {
	ObjectType string      `json:"object_type"`
	Object     interface{} `json:"object"`
}

func init() {
	producer = NewDeltaMetricsMessageProducer()
}

func NewDeltaMetricsTracer() *DeltaMetricsTracer {
	return &DeltaMetricsTracer{
		Producer: producer,
	}
}

func (p *DeltaMetricsTracer) Trace(message DeltaMetricsBaseMessage) error {
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	producer.Publish(messageBytes)
	return nil
}

func (p *DeltaMetricsTracer) TraceMessageString(message string) error {
	err := producer.Publish([]byte(message))
	if err != nil {
		return err
	}
	return nil
}

func (p *DeltaMetricsTracer) TraceMessage(message []byte) error {
	err := producer.Publish(message)
	if err != nil {
		return err
	}
	return nil
}
