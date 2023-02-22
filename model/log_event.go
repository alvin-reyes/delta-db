package model

import (
	"github.com/libp2p/go-libp2p/core/protocol"
	"time"
)

// time series log events

type LogEvent struct {
	ID             int64     `gorm:"primaryKey"` // auto increment
	LogEventType   string    `json:"log_event"`  // content, deal, piece_commitment, upload, miner, info
	LogEventObject string    `json:"event_object"`
	LogEventId     int64     `json:"log_event_id"` // object id
	LogEvent       string    `json:"log_event"`    // description
	CreatedAt      time.Time `json:"created_at"`   // auto set
	UpdatedAt      time.Time `json:"updated_at"`
}

// api
type ContentUploadContentLog struct {
	ID               int64     `gorm:"primaryKey"` // auto increment
	RequestingApiKey string    `json:"requesting_api_key"`
	EventMessage     string    `json:"event_message"`
	ContentId        int64     `json:"content_id"`
	ContentMeta      string    `json:"content_meta"`
	CreatedAt        time.Time `json:"created_at"` // auto set
	UpdatedAt        time.Time `json:"updated_at"`
}

type ContentCommitmentPieceLog struct {
	ID               int64     `gorm:"primaryKey"` // auto increment
	RequestingApiKey string    `json:"requesting_api_key"`
	EventMessage     string    `json:"event_message"`
	ContentId        int64     `json:"content_id"`
	ContentMeta      string    `json:"content_meta"`
	CreatedAt        time.Time `json:"created_at"` // auto set
	UpdatedAt        time.Time `json:"updated_at"`
}

type ContentCommitmentPiecesLog struct {
	ID               int64     `gorm:"primaryKey"` // auto increment
	BatchRequestUUID string    `json:"batch_request_uuid"`
	RequestingApiKey string    `json:"requesting_api_key"`
	EventMessage     string    `json:"event_message"`
	ContentId        int64     `json:"content_id"`
	ContentMeta      string    `json:"content_meta"`
	CreatedAt        time.Time `json:"created_at"` // auto set
	UpdatedAt        time.Time `json:"updated_at"`
}

type ContentProposalLog struct {
}

type ContentAnnounceLog struct {
}

type ContentOpenStatsLog struct {
}

type RepairRequestLog struct {
}

type NodeRequestLog struct {
}

// job events
type PieceCommitmentJobLog struct {
}

type StorageDealMakeJobLog struct {
}

type DataTransferStatusJobLog struct {
}

type InstanceMetaJobLog struct {
}

// time series log events
type ContentLog struct {
	ID               int64     `gorm:"primaryKey"` // auto increment
	RequestingApiKey string    `json:"requesting_api_key"`
	ContentId        int64     `json:"content_id"`
	CreatedAt        time.Time `json:"created_at"` // auto set
	UpdatedAt        time.Time `json:"updated_at"`
}

// time series content deal events
type ContentDealLog struct {
	ID                  int64       `gorm:"primaryKey"` // auto increment
	RequestingApiKey    string      `json:"requesting_api_key"`
	ContentDealId       int64       `json:"content_deal_id"`
	PropCid             string      `json:"propCid"`
	DealUUID            string      `json:"dealUuid"`
	Miner               string      `json:"miner"`
	DealID              int64       `json:"dealId"`
	Failed              bool        `json:"failed"`
	Verified            bool        `json:"verified"`
	Slashed             bool        `json:"slashed"`
	FailedAt            time.Time   `json:"failedAt,omitempty"`
	DTChan              string      `json:"dtChan" gorm:"index"`
	TransferStarted     time.Time   `json:"transferStarted"`
	TransferFinished    time.Time   `json:"transferFinished"`
	OnChainAt           time.Time   `json:"onChainAt"`
	SealedAt            time.Time   `json:"sealedAt"`
	LastMessage         string      `json:"lastMessage"`
	DealProtocolVersion protocol.ID `json:"deal_protocol_version"`
	MinerVersion        string      `json:"miner_version,omitempty"`
	CreatedAt           time.Time   `json:"created_at"` // auto set
	UpdatedAt           time.Time   `json:"updated_at"`
}

type ContentMinerLog struct {
	ID               int64     `gorm:"primaryKey"` // auto increment
	RequestingApiKey string    `json:"requesting_api_key"`
	ContentMinerId   int64     `json:"content_miner_id"`
	CreatedAt        time.Time `json:"created_at"` // auto set
	UpdatedAt        time.Time `json:"updated_at"`
}

type ContentPieceCommitmentLog struct {
	ID                       int64     `gorm:"primaryKey"` // auto increment
	RequestingApiKey         string    `json:"requesting_api_key"`
	ContentPieceCommitmentId int64     `json:"content_piece_commitment_id"`
	RelatedContentId         int64     `json:"related_content_id"`
	CreatedAt                time.Time `json:"created_at"` // auto set
	UpdatedAt                time.Time `json:"updated_at"`
}

type ContentDealProposalLog struct {
	ID                    int64     `gorm:"primaryKey"` // auto increment
	RequestingApiKey      string    `json:"requesting_api_key"`
	ContentDealProposalID int64     `json:"content_deal_proposal_id"`
	RelatedContentId      int64     `json:"related_content_id"`
	CreatedAt             time.Time `json:"created_at"` // auto set
	UpdatedAt             time.Time `json:"updated_at"`
}
