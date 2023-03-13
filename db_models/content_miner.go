package db_models

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type ContentMiner struct {
	ID        int64     `gorm:"primaryKey"`
	Content   int64     `json:"content" gorm:"index:,option:CONCURRENTLY"`
	Miner     string    `json:"miner"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *ContentMiner) AfterSave(tx *gorm.DB) (err error) {

	if u.Miner == "" {
		return
	}

	messageBytes, err := json.Marshal(u)
	tx.Model(&LogEvent{}).Save(&LogEvent{
		LogEventType:   "ContentMiner",
		LogEventObject: messageBytes,
		LogEventId:     u.ID,
		Collected:      false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	})
	return
}
