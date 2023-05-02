package db_models

import (
	"time"
)

// create an entry first
type BatchImport struct {
	ID        int64     `gorm:"primaryKey"`
	Uuid      string    `json:"uuid" gorm:"index:,option:CONCURRENTLY"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// associate the content to a batch
type BatchContent struct {
	ID            int64     `gorm:"primaryKey"`
	BatchImportID int64     `json:"batch_import_id" gorm:"index:,option:CONCURRENTLY"`
	ContentID     int64     `json:"content_id" gorm:"index:,option:CONCURRENTLY"` // check status of the content
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
