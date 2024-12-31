package jobs

import (
	"log"

	"bbb-voting-service/internal/application/usecases/cache"
)

type SyncCacheJob struct {
	SyncCacheUsecase             *cache.SyncResultsCacheUsecase
	SyncParticipantsCacheUsecase *cache.SyncParticipantsCacheUsecase
}

func NewSyncCacheJob(syncCacheUsecase *cache.SyncResultsCacheUsecase, syncParticipantsCacheUsecase *cache.SyncParticipantsCacheUsecase) *SyncCacheJob {
	return &SyncCacheJob{
		SyncCacheUsecase:             syncCacheUsecase,
		SyncParticipantsCacheUsecase: syncParticipantsCacheUsecase,
	}
}

func (job *SyncCacheJob) Run() {
	err := job.SyncCacheUsecase.Execute()
	if err != nil {
		log.Printf("Error synchronizing results cache: %v", err)
	} else {
		log.Println("Results cache synchronized successfully")
	}

	err = job.SyncParticipantsCacheUsecase.Execute()
	if err != nil {
		log.Printf("Error synchronizing participants cache: %v", err)
	} else {
		log.Println("Participants cache synchronized successfully")
	}
}
