package main

import (
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
