package db_models

import (
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
	ID                  int64               `gorm:"primaryKey"` // auto increment
	NodeInfo            string              `json:"node_info"`
	RequestingApiKey    string              `json:"requesting_api_key"`
	EventMessage        string              `json:"event_message"`
	ContentDealProposal ContentDealProposal `json:"content_deal_proposal"`
	CreatedAt           time.Time           `json:"created_at"` // auto set
	UpdatedAt           time.Time           `json:"updated_at"`
}
type ContentAnnounceLog struct {
	ID                  int64               `gorm:"primaryKey"` // auto increment
	NodeInfo            string              `json:"node_info"`
	RequestingApiKey    string              `json:"requesting_api_key"`
	EventMessage        string              `json:"event_message"`
	ContentDealProposal ContentDealProposal `json:"content_deal_proposal"`
	CreatedAt           time.Time           `json:"created_at"` // auto set
	UpdatedAt           time.Time           `json:"updated_at"`
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
	ID                int64     `gorm:"primaryKey"` // auto increment
	Name              string    `json:"name"`
	Size              int64     `json:"size"`
	Cid               string    `json:"cid"`
	RequestingApiKey  string    `json:"requesting_api_key,omitempty"`
	PieceCommitmentId int64     `json:"piece_commitment_id,omitempty"`
	Status            string    `json:"status"`
	ConnectionMode    string    `json:"connection_mode"` // offline or online
	LastMessage       string    `json:"last_message"`
	NodeInfo          string    `json:"node_info"`
	RequesterInfo     string    `json:"requester_info"`
	SystemContentId   int64     `json:"system_content_id"`
	CreatedAt         time.Time `json:"created_at"` // auto set
	UpdatedAt         time.Time `json:"updated_at"`
}

// ContentDealLog time series content deal events
type ContentDealLog struct {
	ID                  int64     `gorm:"primaryKey"`
	Content             int64     `json:"content" gorm:"index:,option:CONCURRENTLY"`
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
	NodeInfo            string    `json:"node_info"`
	RequesterInfo       string    `json:"requester_info"`
	RequestingApiKey    string    `json:"requesting_api_key"`
	SystemContentDealId int64     `json:"system_content_deal_id"`
	CreatedAt           time.Time `json:"created_at"` // auto set
	UpdatedAt           time.Time `json:"updated_at"`
}

type ContentMinerLog struct {
	ID                   int64     `gorm:"primaryKey"`
	Content              int64     `json:"content" gorm:"index:,option:CONCURRENTLY"`
	Miner                string    `json:"miner"`
	NodeInfo             string    `json:"node_info"`
	RequesterInfo        string    `json:"requester_info"`
	RequestingApiKey     string    `json:"requesting_api_key"`
	SystemContentMinerId int64     `json:"system_content_miner_id"`
	CreatedAt            time.Time `json:"created_at"` // auto set
	UpdatedAt            time.Time `json:"updated_at"`
}

type ContentWalletLog struct {
	ID                    int64     `gorm:"primaryKey"`
	Content               int64     `json:"content" gorm:"index:,option:CONCURRENTLY"`
	Wallet                string    `json:"wallet_meta"`
	NodeInfo              string    `json:"node_info"`
	RequesterInfo         string    `json:"requester_info"`
	RequestingApiKey      string    `json:"requesting_api_key"`
	SystemContentWalletId int64     `json:"system_content_miner_id"`
	CreatedAt             time.Time `json:"created_at"` // auto set
	UpdatedAt             time.Time `json:"updated_at"`
}

type PieceCommitmentLog struct {
	ID                             int64     `gorm:"primaryKey"`
	Cid                            string    `json:"cid"`
	Piece                          string    `json:"piece"`
	Size                           int64     `json:"size"`
	PaddedPieceSize                uint64    `json:"padded_piece_size"`
	UnPaddedPieceSize              uint64    `json:"unnpadded_piece_size"`
	Status                         string    `json:"status"` // open, in-progress, completed (closed).
	LastMessage                    string    `json:"last_message"`
	NodeInfo                       string    `json:"node_info"`
	RequesterInfo                  string    `json:"requester_info"`
	RequestingApiKey               string    `json:"requesting_api_key"`
	SystemContentPieceCommitmentId int64     `json:"system_content_piece_commitment_id"`
	CreatedAt                      time.Time `json:"created_at"` // auto set
	UpdatedAt                      time.Time `json:"updated_at"`
}

type ContentDealProposalLog struct {
	ID                          int64     `gorm:"primaryKey"`
	Content                     int64     `json:"content" gorm:"index:,option:CONCURRENTLY"`
	Unsigned                    string    `json:"unsigned"`
	Signed                      string    `json:"signed"`
	Meta                        string    `json:"meta"`
	NodeInfo                    string    `json:"node_info"`
	RequesterInfo               string    `json:"requester_info"`
	RequestingApiKey            string    `json:"requesting_api_key"`
	SystemContentDealProposalId int64     `json:"system_content_deal_proposal_id"`
	CreatedAt                   time.Time `json:"created_at"` // auto set
	UpdatedAt                   time.Time `json:"updated_at"`
}

type ContentDealProposalParametersLog struct {
	ID                                    int64     `gorm:"primaryKey"`
	Content                               int64     `json:"content" gorm:"index:,option:CONCURRENTLY"`
	Label                                 string    `json:"label,omitempty"`
	Duration                              int64     `json:"duration,omitempty"`
	StartEpoch                            int64     `json:"start_epoch,omitempty"`
	EndEpoch                              int64     `json:"end_epoch,omitempty"`
	TransferParams                        string    `json:"transfer_params,omitempty"`
	RemoveUnsealedCopy                    bool      `json:"remove_unsealed_copy,omitempty"`
	SkipIPNIAnnounce                      bool      `json:"skip_ipni_announce,omitempty"`
	NodeInfo                              string    `json:"node_info"`
	RequesterInfo                         string    `json:"requester_info"`
	RequestingApiKey                      string    `json:"requesting_api_key"`
	SystemContentDealProposalParametersId int64     `json:"system_content_deal_proposal_id"`
	CreatedAt                             time.Time `json:"created_at"` // auto set
	UpdatedAt                             time.Time `json:"updated_at"`
}

type WalletLog struct {
	ID               int64     `gorm:"primaryKey"`
	UuId             string    `json:"uuid"`
	Addr             string    `json:"addr"`
	Owner            string    `json:"owner"`
	KeyType          string    `json:"key_type"`
	PrivateKey       string    `json:"private_key"`
	NodeInfo         string    `json:"node_info"`
	RequesterInfo    string    `json:"requester_info"`
	RequestingApiKey string    `json:"requesting_api_key"`
	SystemWalletId   int64     `json:"system_content_deal_proposal_id"`
	CreatedAt        time.Time `json:"created_at"` // auto set
	UpdatedAt        time.Time `json:"updated_at"`
}
