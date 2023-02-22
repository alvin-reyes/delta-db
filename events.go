package events

import (
	"delta-events/model"
	"gorm.io/gorm"
)

type DeltaEventsLog interface {
	LogInformation()
	LogError()
	LogWarning()
	LogDebug()
}

type DeltaEventDB struct {
	LocalDB *gorm.DB
	TimeDB  *gorm.DB
}

var (
	metricsTopicUrl = ""
	primaryTopic    = "delta-metric-events"
)

type DeltaDBParams struct {
	// create instance of the delta db - connect or create local db and connect to timescaledb if enabled.
	LocalDsn                string // local db dsn
	TimeDsn                 string // time db dsn
	EnableMetricsCollection bool
}

func NewDeltaModels(deltaDbParams DeltaDBParams) DeltaEventDB {
	localDb, err := model.OpenLocalDatabase(deltaDbParams.LocalDsn)
	if err != nil {
		panic(err)
	}

	remoteDb, err := model.OpenRemoteDatabase(deltaDbParams.TimeDsn)
	if err != nil {
		panic(err)
	}

	return DeltaEventDB{
		LocalDB: localDb,
		TimeDB:  remoteDb,
	}
}

// allow dev to define their own delta log event types
// allow dev to define their own delta log events
// allow dev to define their own delta metrics
