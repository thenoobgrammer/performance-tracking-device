package queue

import (
	"errors"
	"log"
	"metric-gathering-tool-service/api"

	"github.com/google/uuid"
)

const DefaultPoolSize uint = 10
const SuccessCode = 0x01
const FailedToEnqueueCode = 0x02
const QueuesAreFullCode = 0x03

type Pool struct {
	ID       uuid.UUID
	Queues   []Queue
	PoolSize uint
	Count    uint
}

func NewPool() Pool {
	return Pool{
		ID:       uuid.New(),
		Queues:   make([]Queue, DefaultPoolSize),
		PoolSize: DefaultPoolSize,
		Count:    0,
	}
}

func (p *Pool) Start() {
	if len(p.Queues) > 0 {
		log.Printf("pool already contains queues")
	}
	for queue := range p.Queues {
		queue.
	}
}

func (p *Pool) Add(s Queue) error {
	if p.Count == p.PoolSize {
		return errors.New("main_pool is full")
	}
	p.Queues = append([]Queue{s}, p.Queues...)
	p.Count++
	return nil
}

func (p *Pool) Remove(id uuid.UUID) {
	for i, b := range p.Queues {
		if b.ID == id {
			p.Queues = append(p.Queues[:i], p.Queues[i+1:]...)
			break
		}
	}
}

func (p *Pool) AllocateRessource(data api.RawData) int {
	for _, q := range p.Queues {
		if q.HasSpace() {
			q.Enqueue(data)
			return SuccessCode
		}
	}
	return QueuesAreFullCode
}
