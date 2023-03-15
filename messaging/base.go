package messaging

import (
	"encoding/json"
	"time"
)

var (
	MetricsTopicUrl = "145.40.77.207:4150"
	PrimaryTopic    = "delta-metric-events"
	PrimaryChannel  = "delta-logs"
	producer        *DeltaMetricsMessageProducer
)

type DeltaMetricsTracer struct {
	Producer *DeltaMetricsMessageProducer
}
type DeltaMetricsBaseMessage struct {
	ObjectType string      `json:"object_type"`
	Object     interface{} `json:"object"`
}

type LogEvent struct {
	ID             int64     `gorm:"primaryKey"` // auto increment
	SourceHost     string    `json:"source_host"`
	SourceIP       string    `json:"source_ip"`
	LogEventType   string    `json:"log_event_type"` // content, deal, piece_commitment, upload, miner, info
	LogEventObject []byte    `json:"event_object"`
	LogEventId     int64     `json:"log_event_id"` // object id
	LogEvent       string    `json:"log_event"`    // description
	DeltaUuid      string    `json:"delta_uuid"`
	CreatedAt      time.Time `json:"created_at"` // auto set
	UpdatedAt      time.Time `json:"updated_at"`
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

func (p *DeltaMetricsTracer) TraceLog(message LogEvent) error {
	baseMessage := DeltaMetricsBaseMessage{
		ObjectType: "LogEvent",
		Object:     message,
	}
	messageBytes, err := json.Marshal(baseMessage)
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
