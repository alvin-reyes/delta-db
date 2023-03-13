package db_models

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type ContentWallet struct {
	ID        int64     `gorm:"primaryKey"`
	Content   int64     `json:"content" gorm:"index:,option:CONCURRENTLY"`
	Wallet    string    `json:"wallet_meta"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *ContentWallet) AfterSave(tx *gorm.DB) (err error) {

	if u.Wallet == "" {
		return
	}

	messageBytes, err := json.Marshal(u)
	tx.Model(&LogEvent{}).Save(&LogEvent{
		LogEventType:   "ContentWallet",
		LogEventObject: messageBytes,
		LogEventId:     u.ID,
		Collected:      false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	})
	return
}
