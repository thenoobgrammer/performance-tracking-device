package queue

import (
	"metric-gathering-tool-service/api"
	"time"

	"github.com/google/uuid"
)

const Capacity = 1024

type Queue struct {
	ID          uuid.UUID
	Buffer      chan api.RawData
	Count       uint
	Capacity    uint
	LastUpdated time.Time
}

func NewQueue() Queue {
	id := uuid.New()
	q := Queue{
		ID:          id,
		Buffer:      make(chan api.RawData, Capacity),
		Count:       0,
		Capacity:    Capacity,
		LastUpdated: time.Now(),
	}

	return q
}

func (q *Queue) Enqueue(payload api.RawData) bool {
	if q.Count >= q.Capacity {
		return false
	}

	q.Buffer <- payload
	q.Count++

	return true
}

func (q *Queue) Dequeue() *api.RawData {
	if q.Count == 0 {
		return nil
	}

	rawData := <-q.Buffer
	q.Count--

	return &rawData
}

func (q *Queue) HasSpace() bool {
	return q.Count < q.Capacity
}
