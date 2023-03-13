package db_models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Content struct {
	ID                int64     `gorm:"primaryKey"`
	Name              string    `json:"name"`
	Size              int64     `json:"size"`
	Cid               string    `json:"cid"`
	RequestingApiKey  string    `json:"requesting_api_key,omitempty"`
	PieceCommitmentId int64     `json:"piece_commitment_id,omitempty"`
	Status            string    `json:"status"`
	ConnectionMode    string    `json:"connection_mode"` // offline or online
	LastMessage       string    `json:"last_message"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (u *Content) AfterSave(tx *gorm.DB) (err error) {

	// log this on the event log table
	messageBytes, err := json.Marshal(u)
	tx.Model(&LogEvent{}).Save(&LogEvent{
		LogEventType:   "Content",
		LogEventObject: messageBytes,
		LogEventId:     u.ID,
		LogEvent:       fmt.Sprintf("Content %d saved", u.ID),
		Collected:      false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	})
	return
}

func Transcode(in, out interface{}) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(in)
	json.NewDecoder(buf).Decode(out)
}
