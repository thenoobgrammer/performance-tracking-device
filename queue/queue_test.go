package queue

import (
	"metric-gathering-tool-service"
	"reflect"
	"testing"
)

func TestStreamEnqueue(t *testing.T) {
	s := NewQueue()

	expect := main.RawData{
		DeviceID:     "randomUUID",
		Lat:          48.283832,
		Lng:          47.323213,
		Height:       1.77,
		Gspeed:       1.2,
		VelUp:        1.2,
		VelDown:      1.2,
		VelX:         1.2,
		Timestamp_ms: 34621821,
		X:            22,
		Y:            22,
		Z:            23,
		Hr:           125,
	}

	s.Enqueue(expect)

	if !reflect.DeepEqual(expect, s.Buffer[0]) {
		t.Errorf("Queue.Enqueue: expected %v, got %v", expect, s.Buffer[0])
	}
}

func TestStreamDequeue(t *testing.T) {
	s := NewQueue()

	expect := main.RawData{
		DeviceID:     "randomUUID",
		Lat:          48.283832,
		Lng:          47.323213,
		Height:       1.77,
		Gspeed:       1.2,
		VelUp:        1.2,
		VelDown:      1.2,
		VelX:         1.2,
		Timestamp_ms: 34621821,
		X:            22,
		Y:            22,
		Z:            23,
		Hr:           125,
	}

	s.Enqueue(expect)

	if s.Count != 1 {
		t.Errorf("Queue.Enqueue: expected %d, got %d", s.Count, len(s.Buffer))
	}

	s.Dequeue()

	if s.Count != 0 {
		t.Errorf("Queue.Dequeue: expected %d, got %d", 0, len(s.Buffer))
	}
}
