package scheduler

import (
	"apigocovid/src/domain"

	"github.com/robfig/cron/v3"
)

// StartScheduler is a function to run service automatically within requested time without hit service
func StartScheduler(cv domain.CovidUsecase) {
	c := cron.New()

	// Call Logging; Every 1 Day at 00 O'Clock
	// c.AddFunc("@midnight", middleware.LoggingActivity)

	// Add Master Statistic; Every 1 Day at 00 O'Clock
	c.AddFunc("0 0 * * *", cv.AddMasterStatistic)

	// Testing
	// c.AddFunc("@every 1m", cv.AddMasterStatistic)

	c.Start()
}
