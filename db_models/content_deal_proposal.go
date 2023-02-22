package db_models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type ContentDealProposal struct {
	ID       int64  `gorm:"primaryKey"`
	Content  int64  `json:"content" gorm:"index:,option:CONCURRENTLY"`
	Unsigned string `json:"unsigned"`
	Signed   string `json:"signed"`
	Meta     string `json:"meta"`
}

func (u *ContentDealProposal) BeforeSave(tx *gorm.DB) (err error) {
	tx.Model(&LogEvent{}).Save(&LogEvent{
		LogEventType: "ContentDealProposal Save",
		LogEventId:   u.ID,
		LogEvent:     fmt.Sprintf("ContentDealProposalParameters %d saved", u.ID),
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	})
	return
}

func (u *ContentDealProposal) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Model(&LogEvent{}).Save(&LogEvent{
		LogEventType: "ContentDealProposal Create",
		LogEventId:   u.ID,
		LogEvent:     fmt.Sprintf("ContentDealProposalParameters %d create", u.ID),
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	})
	return
}

func (u *ContentDealProposal) AfterSave(tx *gorm.DB) (err error) {
	tx.Model(&LogEvent{}).Save(&LogEvent{
		LogEventType: "After ContentDealProposalParameters Save",
		LogEventId:   u.ID,
		LogEvent:     fmt.Sprintf("After ContentDealProposalParameters %d saved", u.ID),
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	})
	return
}
