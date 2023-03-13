package db_models

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type ContentDealProposalParameters struct {
	ID                 int64     `gorm:"primaryKey"`
	Content            int64     `json:"content" gorm:"index:,option:CONCURRENTLY"`
	Label              string    `json:"label,omitempty"`
	Duration           int64     `json:"duration,omitempty"`
	StartEpoch         int64     `json:"start_epoch,omitempty"`
	EndEpoch           int64     `json:"end_epoch,omitempty"`
	TransferParams     string    `json:"transfer_params,omitempty"`
	RemoveUnsealedCopy bool      `json:"remove_unsealed_copy,omitempty"`
	SkipIPNIAnnounce   bool      `json:"skip_ipni_announce,omitempty"`
	CreatedAt          time.Time `json:"created_at" json:"created-at"`
	UpdatedAt          time.Time `json:"updated_at" json:"updated-at"`
}

func (u *ContentDealProposalParameters) AfterSave(tx *gorm.DB) (err error) {
	messageBytes, err := json.Marshal(u)
	tx.Model(&LogEvent{}).Save(&LogEvent{
		LogEventType:   "ContentDealProposalParameters",
		LogEventObject: messageBytes,
		LogEventId:     u.ID,
		Collected:      false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	})
	return
}
