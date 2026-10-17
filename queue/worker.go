package queue

import (
	"log"
	"metric-gathering-tool-service/api"
)

func worker(jobs chan api.RawData) {
	for job := range jobs {
		process(job)
	}
}

func process(job api.RawData) {
	log.Printf("RawData: %s", job.DeviceID)
}
