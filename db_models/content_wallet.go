package db_models

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type ContentWallet struct {
	ID        int64     `gorm:"primaryKey"`
	Content   int64     `json:"content" gorm:"index:,option:CONCURRENTLY"`
	Wallet    string    `json:"wallet_meta"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *ContentWallet) AfterCreate(tx *gorm.DB) (err error) {

	var instanceFromDb InstanceMeta
	tx.Raw("SELECT * FROM instance_meta ORDER BY id DESC LIMIT 1").Scan(&instanceFromDb)

	if instanceFromDb.ID == 0 {
		return
	}

	var contentWallet ContentWallet
	tx.Model(&ContentWallet{}).Where("id = ?", u.ID).First(&contentWallet)

	if contentWallet.ID == 0 {
		return
	}

	// get instance info
	ip, err := GetPublicIP()
	if err != nil {
		return
	}
	log := ContentWalletLog{
		Content:               u.Content,
		Wallet:                u.Wallet,
		NodeInfo:              GetHostname(),
		RequesterInfo:         ip,
		SystemContentWalletId: u.ID,
		DeltaNodeUuid:         instanceFromDb.InstanceUuid,
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),
	}

	deltaMetricsBaseMessage := DeltaMetricsBaseMessage{
		ObjectType: "ContentWalletLog",
		Object:     log,
	}

	messageBytes, err := json.Marshal(deltaMetricsBaseMessage)
	producer.Publish(messageBytes)

	return
}
