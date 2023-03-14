package db_models

import (
	"encoding/json"
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

	var instanceFromDb InstanceMeta
	// get the latest instance info based on created_at
	tx.Model(&InstanceMeta{}).Where("id > 0").First(&instanceFromDb)

	if instanceFromDb.ID == 0 {
		return
	}

	var contentFromDb Content
	tx.Model(&Content{}).Where("id = ?", u.ID).First(&contentFromDb)

	if contentFromDb.ID == 0 {
		return
	}
	// get instance info
	ip, err := GetPublicIP()
	if err != nil {
		return
	}

	log := ContentLog{
		Name:              contentFromDb.Name,
		Size:              contentFromDb.Size,
		Cid:               contentFromDb.Cid,
		RequestingApiKey:  contentFromDb.RequestingApiKey,
		PieceCommitmentId: contentFromDb.PieceCommitmentId,
		Status:            contentFromDb.Status,
		ConnectionMode:    contentFromDb.ConnectionMode,
		LastMessage:       contentFromDb.LastMessage,
		NodeInfo:          GetHostname(),
		RequesterInfo:     ip,
		DeltaNodeUuid:     instanceFromDb.InstanceUuid,
		SystemContentId:   contentFromDb.ID,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	deltaMetricsBaseMessage := DeltaMetricsBaseMessage{
		ObjectType: "ContentLog",
		Object:     log,
	}

	messageBytes, err := json.Marshal(deltaMetricsBaseMessage)
	if err != nil {
		return err
	}
	producer.Publish(messageBytes)

	return
}
