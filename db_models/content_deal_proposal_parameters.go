package db_models

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type ContentDealProposalParameters struct {
	ID                 int64     `gorm:"primaryKey"`
	Content            int64     `json:"content" gorm:"index:,option:CONCURRENTLY"`
	Label              string    `json:"label,omitempty"`
	Duration           int64     `json:"duration,omitempty"`
	StartEpoch         int64     `json:"start_epoch,omitempty"`
	EndEpoch           int64     `json:"end_epoch,omitempty"`
	TransferParams     string    `json:"transfer_params,omitempty"`
	RemoveUnsealedCopy bool      `json:"remove_unsealed_copy,omitempty"`
	SkipIPNIAnnounce   bool      `json:"skip_ipni_announce,omitempty"`
	CreatedAt          time.Time `json:"created_at" json:"created-at"`
	UpdatedAt          time.Time `json:"updated_at" json:"updated-at"`
}

func (u *ContentDealProposalParameters) AfterCreate(tx *gorm.DB) (err error) {

	var instanceFromDb InstanceMeta
	tx.Model(&InstanceMeta{}).Where("id > 0").First(&instanceFromDb)

	if instanceFromDb.ID == 0 {
		return
	}

	var contentDealProposalParams ContentDealProposalParameters
	tx.Model(&ContentDealProposalParameters{}).Where("id = ?", u.ID).First(&contentDealProposalParams)

	if contentDealProposalParams.ID == 0 {
		return
	}

	// get instance info
	ip, err := GetPublicIP()
	if err != nil {
		return
	}
	log := ContentDealProposalParametersLog{
		Content:                               contentDealProposalParams.Content,
		Label:                                 contentDealProposalParams.Label,
		Duration:                              contentDealProposalParams.Duration,
		StartEpoch:                            contentDealProposalParams.StartEpoch,
		EndEpoch:                              contentDealProposalParams.EndEpoch,
		TransferParams:                        contentDealProposalParams.TransferParams,
		RemoveUnsealedCopy:                    contentDealProposalParams.RemoveUnsealedCopy,
		SkipIPNIAnnounce:                      contentDealProposalParams.SkipIPNIAnnounce,
		NodeInfo:                              GetHostname(),
		RequesterInfo:                         ip,
		DeltaNodeUuid:                         instanceFromDb.InstanceUuid,
		SystemContentDealProposalParametersId: u.ID,
		CreatedAt:                             time.Now(),
		UpdatedAt:                             time.Now(),
	}

	deltaMetricsBaseMessage := DeltaMetricsBaseMessage{
		ObjectType: "ContentDealProposalParametersLog",
		Object:     log,
	}

	messageBytes, err := json.Marshal(deltaMetricsBaseMessage)
	producer.Publish(messageBytes)

	return
}
