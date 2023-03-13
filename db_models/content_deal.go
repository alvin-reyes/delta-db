package db_models

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type ContentDeal struct {
	ID      int64 `gorm:"primaryKey"`
	Content int64 `json:"content" gorm:"index:,option:CONCURRENTLY"`
	//Content             Content   `gorm:"references:ID"`
	PropCid             string    `json:"propCid"`
	DealUUID            string    `json:"dealUuid"`
	Miner               string    `json:"miner"`
	DealID              int64     `json:"dealId"`
	Failed              bool      `json:"failed"`
	Verified            bool      `json:"verified"`
	Slashed             bool      `json:"slashed"`
	FailedAt            time.Time `json:"failedAt,omitempty"`
	DTChan              string    `json:"dtChan" gorm:"index"`
	TransferStarted     time.Time `json:"transferStarted"`
	TransferFinished    time.Time `json:"transferFinished"`
	OnChainAt           time.Time `json:"onChainAt"`
	SealedAt            time.Time `json:"sealedAt"`
	LastMessage         string    `json:"lastMessage"`
	DealProtocolVersion string    `json:"deal_protocol_version"`
	MinerVersion        string    `json:"miner_version,omitempty"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

func (u *ContentDeal) AfterSave(tx *gorm.DB) (err error) {
	messageBytes, err := json.Marshal(u)
	tx.Model(&LogEvent{}).Save(&LogEvent{
		LogEventType:   "ContentDeal",
		LogEventObject: messageBytes,
		LogEventId:     u.ID,
		Collected:      false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	})
	return
}
