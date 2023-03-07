package event_models

import (
	"github.com/application-research/delta-db/db_models"
	"time"
)

// LogEvent time series log events
// `LogEvent` is a struct with fields `ID`, `LogEventType`, `LogEventObject`, `LogEventId`, `LogEvent`, `CreatedAt`, and
// `UpdatedAt`.
//
// The `gorm:"primaryKey"` tag tells GORM that the `ID` field is the primary key for the table.
//
// The `json:"log_event"` tag tells GORM that the `LogEventType` field should be named `log_event` when it is serialized to
// JSON.
//
// The `json:"created_at"
// @property {int64} ID - The primary key for the table.
// @property {string} LogEventType - The type of event that occurred.
// @property {string} LogEventObject - content, deal, piece_commitment, upload, miner, info
// @property {int64} LogEventId - The id of the object that the event is about.
// @property {string} LogEvent - The type of event that occurred.
// @property CreatedAt - The time the event was created
// @property UpdatedAt - The time the event was created
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
// `ContentUploadContentLog` is a type that has a `Content` field of type `db_models.Content` and a `CreatedAt` field of
// type `time.Time`.
// @property {int64} ID - The primary key for the table.
// @property {string} NodeInfo - The node that is requesting the content upload.
// @property {string} RequestingApiKey - The api key of the node that is requesting the content
// @property {string} EventMessage - The message that will be logged.
// @property Content - The content that was uploaded.
// @property CreatedAt - The time the log was created
// @property UpdatedAt - This is the time the record was last updated.
type ContentUploadContentLog struct {
	ID               int64             `gorm:"primaryKey"` // auto increment
	NodeInfo         string            `json:"node_info"`
	RequestingApiKey string            `json:"requesting_api_key"`
	EventMessage     string            `json:"event_message"`
	Content          db_models.Content `json:"content"`
	CreatedAt        time.Time         `json:"created_at"` // auto set
	UpdatedAt        time.Time         `json:"updated_at"`
}

// `ContentCommitmentPieceLog` is a struct that contains a `Content` and a `PieceCommitment` and a `NodeInfo` and a
// `RequestingApiKey` and a `EventMessage` and a `CreatedAt` and a `UpdatedAt`.
// @property {int64} ID - The primary key for the table.
// @property {string} NodeInfo - The node that is logging the event
// @property {string} RequestingApiKey - The api key of the node that is requesting the content
// @property {string} EventMessage - A string that describes the event that occurred.
// @property Content - The content that was committed to.
// @property ContentPieceCommitment - The piece commitment that was created for the content.
// @property CreatedAt - The time the log was created
// @property UpdatedAt - The time the record was last updated.
// `ContentCommitmentPieceLog` is a struct that contains a `Content` and a `PieceCommitment` and a `NodeInfo` and a
// `RequestingApiKey` and a `EventMessage` and a `CreatedAt` and a `UpdatedAt`.
// @property {int64} ID - The primary key for the table.
// @property {string} NodeInfo - The node that is making the request
// @property {string} RequestingApiKey - The api key of the node that is requesting the content
// @property {string} EventMessage - A string that describes the event that occurred.
// @property Content - The content that was committed to.
// @property ContentPieceCommitment - The piece commitment that was created for the content.
// @property CreatedAt - The time the log was created
// @property UpdatedAt - The time the record was last updated.
type ContentCommitmentPieceLog struct {
	ID                     int64                     `gorm:"primaryKey"` // auto increment
	NodeInfo               string                    `json:"node_info"`
	RequestingApiKey       string                    `json:"requesting_api_key"`
	EventMessage           string                    `json:"event_message"`
	Content                db_models.Content         `json:"content"`
	ContentPieceCommitment db_models.PieceCommitment `json:"piece_commitment"`
	CreatedAt              time.Time                 `json:"created_at"` // auto set
	UpdatedAt              time.Time                 `json:"updated_at"`
}

// `ContentCommitmentPiecesLog` is a struct that contains a `Content` and a slice of `PieceCommitment`s.
// @property {int64} ID - auto incrementing primary key
// @property {string} NodeInfo - The node that is logging the event
// @property {string} BatchRequestUUID - The UUID of the batch request that this content commitment was part of.
// @property {string} RequestingApiKey - The api key of the node that made the request
// @property {string} EventMessage - The message that was sent to the event bus
// @property Content - The content that was committed to.
// @property {[]db_models.PieceCommitment} ContentPieceCommitments - This is a slice of PieceCommitment structs.
// @property CreatedAt - The time the event was created
// @property UpdatedAt - The time the record was last updated.
type ContentCommitmentPiecesLog struct {
	ID                      int64                       `gorm:"primaryKey"` // auto increment
	NodeInfo                string                      `json:"node_info"`
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
	NodeInfo            string                        `json:"node_info"`
	RequestingApiKey    string                        `json:"requesting_api_key"`
	EventMessage        string                        `json:"event_message"`
	ContentDealProposal db_models.ContentDealProposal `json:"content_deal_proposal"`
	CreatedAt           time.Time                     `json:"created_at"` // auto set
	UpdatedAt           time.Time                     `json:"updated_at"`
}

type ContentDealProposalParameterLog struct {
	ID                              int64     `gorm:"primaryKey"` // auto increment
	NodeInfo                        string    `json:"node_info"`
	RequestingApiKey                string    `json:"requesting_api_key"`
	ContentDealProposalParameterLog int64     `json:"content_deal_proposal_id"`
	RelatedContentId                int64     `json:"related_content_id"`
	CreatedAt                       time.Time `json:"created_at"` // auto set
	UpdatedAt                       time.Time `json:"updated_at"`
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
	RequestingApiKey string    `json:"requesting_api_key"`
	EventMessage     string    `json:"event_message"`
	CreatedAt        time.Time `json:"created_at"` // auto set
	UpdatedAt        time.Time `json:"updated_at"`
}

type NodeRequestLog struct {
	ID            int64     `gorm:"primaryKey"` // auto increment
	NodeInfo      string    `json:"node_info"`
	RequesterInfo string    `json:"requester_info"`
	CreatedAt     time.Time `json:"created_at"` // auto set
	UpdatedAt     time.Time `json:"updated_at"`
}

// job events
type PieceCommitmentJobLog struct {
	ID       int64
	NodeInfo string `json:"node_info"`
}

type StorageDealMakeJobLog struct {
	ID       int64
	NodeInfo string `json:"node_info"`
}

type DataTransferStatusJobLog struct {
	ID       int64
	NodeInfo string `json:"node_info"`
}

type InstanceMetaJobLog struct {
	ID       int64
	NodeInfo string `json:"node_info"`
}

// ContentLog time series log events
type ContentLog struct {
	ID               int64     `gorm:"primaryKey"` // auto increment
	NodeInfo         string    `json:"node_info"`
	RequestingApiKey string    `json:"requesting_api_key"`
	ContentId        int64     `json:"content_id"`
	CreatedAt        time.Time `json:"created_at"` // auto set
	UpdatedAt        time.Time `json:"updated_at"`
}

// ContentDealLog time series content deal events
type ContentDealLog struct {
	ID               int64                 `gorm:"primaryKey"` // auto increment
	NodeInfo         string                `json:"node_info"`
	RequestingApiKey string                `json:"requesting_api_key"`
	ContentDealId    int64                 `json:"content_deal_id"`
	ContentDeal      db_models.ContentDeal `json:"content_deal"`
	CreatedAt        time.Time             `json:"created_at"` // auto set
	UpdatedAt        time.Time             `json:"updated_at"`
}

type ContentMinerLog struct {
	ID               int64     `gorm:"primaryKey"` // auto increment
	NodeInfo         string    `json:"node_info"`
	RequestingApiKey string    `json:"requesting_api_key"`
	ContentMinerId   int64     `json:"content_miner_id"`
	CreatedAt        time.Time `json:"created_at"` // auto set
	UpdatedAt        time.Time `json:"updated_at"`
}

type ContentPieceCommitmentLog struct {
	ID                       int64     `gorm:"primaryKey"` // auto increment
	NodeInfo                 string    `json:"node_info"`
	RequestingApiKey         string    `json:"requesting_api_key"`
	ContentPieceCommitmentId int64     `json:"content_piece_commitment_id"`
	RelatedContentId         int64     `json:"related_content_id"`
	CreatedAt                time.Time `json:"created_at"` // auto set
	UpdatedAt                time.Time `json:"updated_at"`
}

type ContentDealProposalLog struct {
	ID                    int64     `gorm:"primaryKey"` // auto increment
	NodeInfo              string    `json:"node_info"`
	RequestingApiKey      string    `json:"requesting_api_key"`
	ContentDealProposalID int64     `json:"content_deal_proposal_id"`
	RelatedContentId      int64     `json:"related_content_id"`
	CreatedAt             time.Time `json:"created_at"` // auto set
	UpdatedAt             time.Time `json:"updated_at"`
}
