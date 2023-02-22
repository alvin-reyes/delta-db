package main

import (
	"github.com/application-research/delta-db/db_models"
	"gorm.io/gorm"
)

type DeltaEventDB struct {
	LocalDB                 *gorm.DB
	TimeDB                  *gorm.DB
	EnableMetricsCollection bool
}

type DeltaEventsLogProducer interface {
	LogInformation()
	LogError()
	LogWarning()
	LogDebug()
}

var (
	metricsTopicUrl = "nsq-test.estuary.tech:4150"
	primaryTopic    = "delta-metric-events"
)

type DeltaDBParams struct {
	// create instance of the delta db - connect or create local db and connect to timescaledb if enabled.
	LocalDbDsn              string // local db dsn
	TimeDbDsn               string // time db dsn
	EnableMetricsCollection bool
}

func NewDeltaModels(deltaDbParams DeltaDBParams) DeltaEventDB {
	localDb, err := db_models.OpenLocalDatabase(deltaDbParams.LocalDbDsn)
	if err != nil {
		panic(err)
	}

	var remoteDb *gorm.DB
	if deltaDbParams.EnableMetricsCollection == true {
		remoteDb, err = db_models.OpenMetricsCollectionDB(deltaDbParams.TimeDbDsn)
		if err != nil {
			panic(err)
		}
	}

	return DeltaEventDB{
		LocalDB:                 localDb,
		TimeDB:                  remoteDb,
		EnableMetricsCollection: deltaDbParams.EnableMetricsCollection,
	}
}
