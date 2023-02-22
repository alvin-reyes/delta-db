package events

import (
	"delta-db/db_models"
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
	metricsTopicUrl = ""
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

// allow dev to define their own delta log event types
// allow dev to define their own delta log events
// allow dev to define their own delta metrics

func (deltaEventDB DeltaEventDB) LogInformation() {

}

//func produceMe() {
//	config := nsq.NewConfig()
//	w, _ := nsq.NewProducer("127.0.0.1:4150", config)
//	for i := 0; i < 1000; i++ {
//		err := w.Publish("eMumba", []byte("test"+strconv.Itoa(i)))
//		if err != nil {
//			log.Panic("Could not connect")
//		}
//	}
//	w.Stop()
//}
