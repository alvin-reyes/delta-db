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

func (u *ContentDeal) AfterCreate(tx *gorm.DB) (err error) {

	var instanceFromDb InstanceMeta
	tx.Model(&InstanceMeta{}).Where("id = ?", 1).First(&instanceFromDb)

	if instanceFromDb.ID == 0 {
		return
	}

	var contentDealLog ContentDeal
	tx.Model(&ContentDeal{}).Where("id = ?", u.ID).First(&contentDealLog)

	if contentDealLog.ID == 0 {
		return
	}
	// get instance info
	ip, err := GetPublicIP()
	if err != nil {
		return
	}

	log := ContentDealLog{
		Content:             contentDealLog.Content,
		PropCid:             contentDealLog.PropCid,
		DealUUID:            contentDealLog.DealUUID,
		Miner:               contentDealLog.Miner,
		DealID:              contentDealLog.DealID,
		Failed:              contentDealLog.Failed,
		Verified:            contentDealLog.Verified,
		Slashed:             contentDealLog.Slashed,
		FailedAt:            contentDealLog.FailedAt,
		DTChan:              contentDealLog.DTChan,
		TransferStarted:     contentDealLog.TransferStarted,
		TransferFinished:    contentDealLog.TransferFinished,
		OnChainAt:           contentDealLog.OnChainAt,
		SealedAt:            contentDealLog.SealedAt,
		LastMessage:         contentDealLog.LastMessage,
		DealProtocolVersion: contentDealLog.DealProtocolVersion,
		MinerVersion:        contentDealLog.MinerVersion,
		DeltaNodeUuid:       instanceFromDb.InstanceUuid,
		NodeInfo:            GetHostname(),
		RequesterInfo:       ip,
		SystemContentDealId: contentDealLog.ID,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}

	deltaMetricsBaseMessage := DeltaMetricsBaseMessage{
		ObjectType: "ContentDealLog",
		Object:     log,
	}

	messageBytes, err := json.Marshal(deltaMetricsBaseMessage)
	if err != nil {
		return err
	}
	producer.Publish(messageBytes)
	return
}
