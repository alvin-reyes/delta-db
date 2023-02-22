package messaging

var (
	MetricsTopicUrl = "nsq-test.estuary.tech:4150"
	PrimaryTopic    = "delta-metric-events"
)

type DeltaMetricsBaseMessage struct {
	ObjectType string      `json:"object_type"`
	Object     interface{} `json:"object"`
}
