package db_models

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type ContentDealProposal struct {
	ID        int64     `gorm:"primaryKey"`
	Content   int64     `json:"content" gorm:"index:,option:CONCURRENTLY"`
	Unsigned  string    `json:"unsigned"`
	Signed    string    `json:"signed"`
	Meta      string    `json:"meta"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *ContentDealProposal) AfterCreate(tx *gorm.DB) (err error) {

	var contentDealProposal ContentDealProposal
	tx.Model(&ContentDealProposal{}).Where("id = ?", u.ID).First(&contentDealProposal)

	if contentDealProposal.ID == 0 {
		return
	}

	// get instance info
	ip, err := GetPublicIP()
	if err != nil {
		return
	}

	log := ContentDealProposalLog{
		Content:                     contentDealProposal.Content,
		Unsigned:                    contentDealProposal.Unsigned,
		Signed:                      contentDealProposal.Signed,
		Meta:                        contentDealProposal.Meta,
		NodeInfo:                    GetHostname(),
		RequesterInfo:               ip,
		SystemContentDealProposalId: u.ID,
		CreatedAt:                   time.Now(),
		UpdatedAt:                   time.Now(),
	}

	deltaMetricsBaseMessage := DeltaMetricsBaseMessage{
		ObjectType: "ContentDealProposalLog",
		Object:     log,
	}

	messageBytes, err := json.Marshal(deltaMetricsBaseMessage)
	producer.Publish(messageBytes)

	return
}
