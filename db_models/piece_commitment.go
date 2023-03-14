package db_models

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type PieceCommitment struct {
	ID                int64     `gorm:"primaryKey"`
	Cid               string    `json:"cid"`
	Piece             string    `json:"piece"`
	Size              int64     `json:"size"`
	PaddedPieceSize   uint64    `json:"padded_piece_size"`
	UnPaddedPieceSize uint64    `json:"unnpadded_piece_size"`
	Status            string    `json:"status"` // open, in-progress, completed (closed).
	LastMessage       string    `json:"last_message"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (u *PieceCommitment) AfterCreate(tx *gorm.DB) (err error) {

	var pieceComm PieceCommitment
	tx.Model(&PieceCommitment{}).Where("id = ?", u.ID).First(&pieceComm)

	if pieceComm.ID == 0 {
		return
	}

	// get instance info
	ip, err := GetPublicIP()
	if err != nil {
		return
	}
	log := PieceCommitmentLog{
		Cid:                            pieceComm.Cid,
		Piece:                          pieceComm.Piece,
		Size:                           pieceComm.Size,
		PaddedPieceSize:                pieceComm.PaddedPieceSize,
		UnPaddedPieceSize:              pieceComm.UnPaddedPieceSize,
		Status:                         pieceComm.Status,
		LastMessage:                    pieceComm.LastMessage,
		NodeInfo:                       GetHostname(),
		RequesterInfo:                  ip,
		SystemContentPieceCommitmentId: u.ID,
		CreatedAt:                      time.Now(),
		UpdatedAt:                      time.Now(),
	}

	deltaMetricsBaseMessage := DeltaMetricsBaseMessage{
		ObjectType: "PieceCommitmentLog",
		Object:     log,
	}

	messageBytes, err := json.Marshal(deltaMetricsBaseMessage)
	producer.Publish(messageBytes)

	return
}
