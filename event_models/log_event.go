package event_models

import (
	"github.com/application-research/delta-db/db_models"
	"time"
)

//
//type LogEvent struct {
//	ID             int64     `gorm:"primaryKey"` // auto increment
//	LogEventType   string    `json:"log_event"`  // content, deal, piece_commitment, upload, miner, info
//	LogEventObject string    `json:"event_object"`
//	LogEventId     int64     `json:"log_event_id"` // object id
//	LogEvent       string    `json:"log_event"`    // description
//	CreatedAt      time.Time `json:"created_at"`   // auto set
//	UpdatedAt      time.Time `json:"updated_at"`
//}

// action events
type DeltaStartupLogs struct {
	ID        int64     `gorm:"primaryKey"` // auto increment
	NodeInfo  string    `json:"node_info"`
	OSDetails string    `json:"os_details"`
	IPAddress string    `json:"ip_address"`
	CreatedAt time.Time `json:"created_at"` // auto set
	UpdatedAt time.Time `json:"updated_at"`
}

type DealEndpointRequestLog struct {
	ID          int64  `gorm:"primaryKey"` // auto increment
	NodeInfo    string `json:"node_info"`
	Information string `json:"information"`
}

type DealContentRequestLog struct {
	ID               int64     `gorm:"primaryKey"` // auto increment
	NodeInfo         string    `json:"node_info"`
	RequestingApiKey string    `json:"requesting_api_key"`
	EventMessage     string    `json:"event_message"`
	CreatedAt        time.Time `json:"created_at"` // auto set
	UpdatedAt        time.Time `json:"updated_at"`
}

type DealPieceCommitmentRequestLog struct {
	ID               int64     `gorm:"primaryKey"` // auto increment
	NodeInfo         string    `json:"node_info"`
	RequestingApiKey string    `json:"requesting_api_key"`
	EventMessage     string    `json:"event_message"`
	CreatedAt        time.Time `json:"created_at"` // auto set
	UpdatedAt        time.Time `json:"updated_at"`
}

type ContentPrepareLog struct {
	ID                  int64                         `gorm:"primaryKey"` // auto increment
	NodeInfo            string                        `json:"node_info"`
	RequestingApiKey    string                        `json:"requesting_api_key"`
	EventMessage        string                        `json:"event_message"`
	ContentDealProposal db_models.ContentDealProposal `json:"content_deal_proposal"`
	CreatedAt           time.Time                     `json:"created_at"` // auto set
	UpdatedAt           time.Time                     `json:"updated_at"`
}
type ContentAnnounceLog struct {
	ID                  int64                         `gorm:"primaryKey"` // auto increment
	NodeInfo            string                        `json:"node_info"`
	RequestingApiKey    string                        `json:"requesting_api_key"`
	EventMessage        string                        `json:"event_message"`
	ContentDealProposal db_models.ContentDealProposal `json:"content_deal_proposal"`
	CreatedAt           time.Time                     `json:"created_at"` // auto set
	UpdatedAt           time.Time                     `json:"updated_at"`
}

type OpenStatsLog struct {
	ID            int64     `gorm:"primaryKey"` // auto increment
	NodeInfo      string    `json:"node_info"`
	RequesterInfo string    `json:"requester_info"`
	CreatedAt     time.Time `json:"created_at"` // auto set
	UpdatedAt     time.Time `json:"updated_at"`
}

type RepairRequestLog struct {
	ID               int64     `gorm:"primaryKey"` // auto increment
	NodeInfo         string    `json:"node_info"`
	RequesterInfo    string    `json:"requester_info"`
	RequestingApiKey string    `json:"requesting_api_key"`
	EventMessage     string    `json:"event_message"`
	CreatedAt        time.Time `json:"created_at"` // auto set
	UpdatedAt        time.Time `json:"updated_at"`
}

type NodeRequestLog struct {
	ID               int64     `gorm:"primaryKey"` // auto increment
	NodeInfo         string    `json:"node_info"`
	RequesterInfo    string    `json:"requester_info"`
	RequestingApiKey string    `json:"requesting_api_key"`
	CreatedAt        time.Time `json:"created_at"` // auto set
	UpdatedAt        time.Time `json:"updated_at"`
}

// job events
type PieceCommitmentJobLog struct {
	ID               int64     `gorm:"primaryKey"` // auto increment
	NodeInfo         string    `json:"node_info"`
	RequesterInfo    string    `json:"requester_info"`
	RequestingApiKey string    `json:"requesting_api_key"`
	CreatedAt        time.Time `json:"created_at"` // auto set
	UpdatedAt        time.Time `json:"updated_at"`
}

type StorageDealMakeJobLog struct {
	ID               int64     `gorm:"primaryKey"` // auto increment
	NodeInfo         string    `json:"node_info"`
	RequesterInfo    string    `json:"requester_info"`
	RequestingApiKey string    `json:"requesting_api_key"`
	CreatedAt        time.Time `json:"created_at"` // auto set
	UpdatedAt        time.Time `json:"updated_at"`
}

type DataTransferStatusJobLog struct {
	ID               int64     `gorm:"primaryKey"` // auto increment
	NodeInfo         string    `json:"node_info"`
	RequesterInfo    string    `json:"requester_info"`
	RequestingApiKey string    `json:"requesting_api_key"`
	CreatedAt        time.Time `json:"created_at"` // auto set
	UpdatedAt        time.Time `json:"updated_at"`
}

type InstanceMetaJobLog struct {
	ID               int64
	NodeInfo         string    `json:"node_info"`
	RequesterInfo    string    `json:"requester_info"`
	RequestingApiKey string    `json:"requesting_api_key"`
	CreatedAt        time.Time `json:"created_at"` // auto set
	UpdatedAt        time.Time `json:"updated_at"`
}

// ContentLog time series log events
type ContentLog struct {
	db_models.Content
	NodeInfo        string `json:"node_info"`
	RequesterInfo   string `json:"requester_info"`
	SystemContentId int64  `json:"system_content_id"`
}

// ContentDealLog time series content deal events
type ContentDealLog struct {
	db_models.ContentDeal
	NodeInfo            string `json:"node_info"`
	RequesterInfo       string `json:"requester_info"`
	RequestingApiKey    string `json:"requesting_api_key"`
	SystemContentDealId int64  `json:"system_content_deal_id"`
}

type ContentMinerLog struct {
	db_models.ContentMiner
	NodeInfo             string `json:"node_info"`
	RequesterInfo        string `json:"requester_info"`
	RequestingApiKey     string `json:"requesting_api_key"`
	SystemContentMinerId int64  `json:"system_content_miner_id"`
}

type PieceCommitmentLog struct {
	db_models.PieceCommitment
	ID                             int64  `gorm:"primaryKey"` // auto increment
	NodeInfo                       string `json:"node_info"`
	RequesterInfo                  string `json:"requester_info"`
	RequestingApiKey               string `json:"requesting_api_key"`
	SystemContentPieceCommitmentId int64  `json:"system_content_piece_commitment_id"`
}

type ContentDealProposalLog struct {
	db_models.ContentDealProposal
	ID                          int64  `gorm:"primaryKey"` // auto increment
	NodeInfo                    string `json:"node_info"`
	RequesterInfo               string `json:"requester_info"`
	RequestingApiKey            string `json:"requesting_api_key"`
	SystemContentDealProposalId int64  `json:"system_content_deal_proposal_id"`
}

type ContentDealProposalParameter struct {
	db_models.ContentDealProposalParameters
	ID                          int64  `gorm:"primaryKey"` // auto increment
	NodeInfo                    string `json:"node_info"`
	RequesterInfo               string `json:"requester_info"`
	RequestingApiKey            string `json:"requesting_api_key"`
	SystemContentDealProposalId int64  `json:"system_content_deal_proposal_id"`
}
