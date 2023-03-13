package db_models

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type ContentDealProposal struct {
	ID        int64     `gorm:"primaryKey"`
	Content   int64     `json:"content" gorm:"index:,option:CONCURRENTLY"`
	Unsigned  string    `json:"unsigned"`
	Signed    string    `json:"signed"`
	Meta      string    `json:"meta"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *ContentDealProposal) AfterSave(tx *gorm.DB) (err error) {
	messageBytes, err := json.Marshal(u)
	tx.Model(&LogEvent{}).Save(&LogEvent{
		LogEventType:   "ContentDealProposal",
		LogEventObject: messageBytes,
		LogEventId:     u.ID,
		Collected:      false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	})
	return
}
