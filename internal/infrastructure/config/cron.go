package config

import (
	"log"
	"os"
	"strconv"

	"bbb-voting-service/internal/infrastructure/jobs"

	"github.com/robfig/cron/v3"
)

func InitCron(syncCacheJob *jobs.SyncCacheJob) *cron.Cron {
	c := cron.New()

	intervalStr := os.Getenv("SYNC_CACHE_INTERVAL")
	if intervalStr == "" {
		intervalStr = "5" // Default to 5 minutes if not set
	}

	_, err := strconv.Atoi(intervalStr)
	if err != nil {
		log.Fatalf("Invalid SYNC_CACHE_INTERVAL value: %v", err)
	}

	// Schedule the job to run every interval minutes
	_, err = c.AddJob("@every "+intervalStr+"m", syncCacheJob)
	if err != nil {
		log.Fatalf("Failed to schedule sync job: %v", err)
	}

	// Execute the job immediately on startup
	go syncCacheJob.Run()

	c.Start()
	return c
}
