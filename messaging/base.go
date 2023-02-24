package messaging

var (
	MetricsTopicUrl = "145.40.77.207:4150"
	PrimaryTopic    = "delta-metric-events"
)

type DeltaMetricsBaseMessage struct {
	ObjectType string      `json:"object_type"`
	Object     interface{} `json:"object"`
}
