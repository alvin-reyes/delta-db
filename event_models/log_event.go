package db_models

import (
	"github.com/application-research/delta-db/db_models"
	"time"
)

// LogEvent time series log events
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
	ID               int64             `gorm:"primaryKey"` // auto increment
	RequestingApiKey string            `json:"requesting_api_key"`
	EventMessage     string            `json:"event_message"`
	Content          db_models.Content `json:"content"`
	CreatedAt        time.Time         `json:"created_at"` // auto set
	UpdatedAt        time.Time         `json:"updated_at"`
}

type ContentCommitmentPieceLog struct {
	ID                     int64                     `gorm:"primaryKey"` // auto increment
	RequestingApiKey       string                    `json:"requesting_api_key"`
	EventMessage           string                    `json:"event_message"`
	Content                db_models.Content         `json:"content"`
	ContentPieceCommitment db_models.PieceCommitment `json:"piece_commitment"`
	CreatedAt              time.Time                 `json:"created_at"` // auto set
	UpdatedAt              time.Time                 `json:"updated_at"`
}

type ContentCommitmentPiecesLog struct {
	ID                      int64                       `gorm:"primaryKey"` // auto increment
	BatchRequestUUID        string                      `json:"batch_request_uuid"`
	RequestingApiKey        string                      `json:"requesting_api_key"`
	EventMessage            string                      `json:"event_message"`
	Content                 db_models.Content           `json:"content"`
	ContentPieceCommitments []db_models.PieceCommitment `json:"piece_commitment"`
	CreatedAt               time.Time                   `json:"created_at"` // auto set
	UpdatedAt               time.Time                   `json:"updated_at"`
}

type ContentProposalLog struct {
	ID                  int64                         `gorm:"primaryKey"` // auto increment
	RequestingApiKey    string                        `json:"requesting_api_key"`
	EventMessage        string                        `json:"event_message"`
	ContentDealProposal db_models.ContentDealProposal `json:"content_deal_proposal"`
	CreatedAt           time.Time                     `json:"created_at"` // auto set
	UpdatedAt           time.Time                     `json:"updated_at"`
}

type ContentAnnounceLog struct {
	ID                  int64                         `gorm:"primaryKey"` // auto increment
	RequestingApiKey    string                        `json:"requesting_api_key"`
	EventMessage        string                        `json:"event_message"`
	ContentDealProposal db_models.ContentDealProposal `json:"content_deal_proposal"`
	CreatedAt           time.Time                     `json:"created_at"` // auto set
	UpdatedAt           time.Time                     `json:"updated_at"`
}

type OpenStatsLog struct {
	ID            int64     `gorm:"primaryKey"` // auto increment
	RequesterInfo string    `json:"requester_info"`
	CreatedAt     time.Time `json:"created_at"` // auto set
	UpdatedAt     time.Time `json:"updated_at"`
}

type RepairRequestLog struct {
	ID               int64     `gorm:"primaryKey"` // auto increment
	RequestingApiKey string    `json:"requesting_api_key"`
	EventMessage     string    `json:"event_message"`
	CreatedAt        time.Time `json:"created_at"` // auto set
	UpdatedAt        time.Time `json:"updated_at"`
}

type NodeRequestLog struct {
	ID            int64     `gorm:"primaryKey"` // auto increment
	RequesterInfo string    `json:"requester_info"`
	CreatedAt     time.Time `json:"created_at"` // auto set
	UpdatedAt     time.Time `json:"updated_at"`
}

// job events
type PieceCommitmentJobLog struct {
	ID int64
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
	ID               int64                 `gorm:"primaryKey"` // auto increment
	RequestingApiKey string                `json:"requesting_api_key"`
	ContentDealId    int64                 `json:"content_deal_id"`
	ContentDeal      db_models.ContentDeal `json:"content_deal"`
	CreatedAt        time.Time             `json:"created_at"` // auto set
	UpdatedAt        time.Time             `json:"updated_at"`
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
