package db_models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type RepairRequest struct {
	ID        int64     `gorm:"primaryKey"`
	ObjectId  int64     `json:"object_id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *RepairRequest) BeforeSave(tx *gorm.DB) (err error) {
	tx.Model(&LogEvent{}).Save(&LogEvent{
		LogEventType: "Content Save",
		LogEventId:   u.ID,
		LogEvent:     fmt.Sprintf("Content %d saved", u.ID),
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	})
	return
}

func (u *RepairRequest) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Model(&LogEvent{}).Save(&LogEvent{
		LogEventType: "Content Create",
		LogEventId:   u.ID,
		LogEvent:     fmt.Sprintf("Content %d create", u.ID),
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	})
	return
}

func (u *RepairRequest) AfterSave(tx *gorm.DB) (err error) {
	tx.Model(&LogEvent{}).Save(&LogEvent{
		LogEventType: "After Content Save",
		LogEventId:   u.ID,
		LogEvent:     fmt.Sprintf("After content %d saved", u.ID),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})
	return
}
