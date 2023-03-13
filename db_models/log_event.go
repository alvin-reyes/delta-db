package db_models

import (
	"encoding/json"
	"fmt"
	"github.com/application-research/delta-db/messaging"
	"gorm.io/gorm"
	"time"
)

var producer *messaging.DeltaMetricsMessageProducer

type DeltaMetricsBaseMessage struct {
	ObjectType string      `json:"object_type"`
	Object     interface{} `json:"object"`
}

func init() {
	producer = messaging.NewDeltaMetricsMessageProducer()
}

// time series log events
type LogEvent struct {
	ID             int64     `gorm:"primaryKey"`     // auto increment
	LogEventType   string    `json:"log_event_type"` // content, deal, piece_commitment, upload, miner, info
	LogEventObject []byte    `json:"event_object"`
	LogEventId     int64     `json:"log_event_id"` // object id
	LogEvent       string    `json:"log_event"`    // description
	Collected      bool      `json:"collected"`    // has this event been collected by the collector?
	CreatedAt      time.Time `json:"created_at"`   // auto set
	UpdatedAt      time.Time `json:"updated_at"`
}

func (u *LogEvent) AfterCreate(tx *gorm.DB) (err error) {
	//	send to collector..
	fmt.Println("LogEvent AfterSave -- time to send it to the collector!!")
	deltaMetricsBaseMessage := DeltaMetricsBaseMessage{
		ObjectType: "LogEvent",
		Object:     u,
	}
	messageBytes, err := json.Marshal(deltaMetricsBaseMessage)
	producer.Publish(messageBytes)
	// update the log event to say it has been collected

	return
}
