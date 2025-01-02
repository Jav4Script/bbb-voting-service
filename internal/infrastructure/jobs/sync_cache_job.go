package jobs

import (
	"context"
	"log"

	"bbb-voting-service/internal/application/usecases/cache"
)

type SyncCacheJob struct {
	SyncResultsCacheUsecase      *cache.SyncResultsCacheUsecase
	SyncParticipantsCacheUsecase *cache.SyncParticipantsCacheUsecase
}

func NewSyncCacheJob(syncCacheUsecase *cache.SyncResultsCacheUsecase, syncParticipantsCacheUsecase *cache.SyncParticipantsCacheUsecase) *SyncCacheJob {
	return &SyncCacheJob{
		SyncResultsCacheUsecase:      syncCacheUsecase,
		SyncParticipantsCacheUsecase: syncParticipantsCacheUsecase,
	}
}

func (job *SyncCacheJob) Run() {
	context := context.Background()

	err := job.SyncResultsCacheUsecase.Execute(context)
	if err != nil {
		log.Printf("Error synchronizing results cache: %v", err)
	} else {
		log.Println("Results cache synchronized successfully")
	}

	err = job.SyncParticipantsCacheUsecase.Execute(context)
	if err != nil {
		log.Printf("Error synchronizing participants cache: %v", err)
	} else {
		log.Println("Participants cache synchronized successfully")
	}
}
