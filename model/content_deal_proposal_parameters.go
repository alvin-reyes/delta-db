package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type ContentDealProposalParameters struct {
	ID             int64     `gorm:"primaryKey"`
	Content        int64     `json:"content" gorm:"index:,option:CONCURRENTLY"`
	Label          string    `json:"label,omitempty"`
	Duration       int64     `json:"duration,omitempty"`
	StartEpoch     int64     `json:"start-epoch,omitempty"`
	EndEpoch       int64     `json:"end-epoch,omitempty"`
	TransferParams string    `json:"transfer-params,omitempty"`
	CreatedAt      time.Time `json:"created_at" json:"created-at"`
	UpdatedAt      time.Time `json:"updated_at" json:"updated-at"`
}

func (u *ContentDealProposalParameters) BeforeSave(tx *gorm.DB) (err error) {
	tx.Model(&LogEvent{}).Save(&LogEvent{
		LogEventType: "ContentDealProposalParameters Save",
		LogEventId:   u.ID,
		LogEvent:     fmt.Sprintf("ContentDealProposalParameters %d saved", u.ID),
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	})
	return
}

func (u *ContentDealProposalParameters) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Model(&LogEvent{}).Save(&LogEvent{
		LogEventType: "ContentDealProposalParameters Create",
		LogEventId:   u.ID,
		LogEvent:     fmt.Sprintf("ContentDealProposalParameters %d create", u.ID),
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	})
	return
}

func (u *ContentDealProposalParameters) AfterSave(tx *gorm.DB) (err error) {
	tx.Model(&LogEvent{}).Save(&LogEvent{
		LogEventType: "After ContentDealProposalParameters Save",
		LogEventId:   u.ID,
		LogEvent:     fmt.Sprintf("After ContentDealProposalParameters %d saved", u.ID),
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	})
	return
}
