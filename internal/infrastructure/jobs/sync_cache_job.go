package jobs

import (
	"log"

	"bbb-voting-service/internal/application/usecases/cache"
)

type SyncCacheJob struct {
	SyncCacheUsecase *cache.SyncCacheUsecase
}

func NewSyncCacheJob(syncCacheUsecase *cache.SyncCacheUsecase) *SyncCacheJob {
	return &SyncCacheJob{
		SyncCacheUsecase: syncCacheUsecase,
	}
}

func (job *SyncCacheJob) Run() {
	err := job.SyncCacheUsecase.Execute()

	if err != nil {
		log.Printf("Error synchronizing cache: %v", err)
	} else {
		log.Println("Cache synchronized successfully")
	}
}
