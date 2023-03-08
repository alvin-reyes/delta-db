package db_models

import (
	"time"
)

// time series log events
type LogEvent struct {
	ID             int64     `gorm:"primaryKey"` // auto increment
	LogEventType   string    `json:"log_event"`  // content, deal, piece_commitment, upload, miner, info
	LogEventObject string    `json:"event_object"`
	LogEventId     int64     `json:"log_event_id"` // object id
	LogEvent       string    `json:"log_event"`    // description
	CreatedAt      time.Time `json:"created_at"`   // auto set
	UpdatedAt      time.Time `json:"updated_at"`
}
