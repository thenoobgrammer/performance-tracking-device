package api

import (
	"encoding/json"
	"net/http"
)

type RawData struct {
	DeviceID     string
	Lat          float32
	Lng          float32
	Height       float32
	Gspeed       float32
	VelUp        float32
	VelDown      float32
	VelX         float32
	Timestamp_ms uint32
	X            int16
	Y            int16
	Z            int16
	Hr           uint8
}

func send(w http.ResponseWriter, r *http.Request) {
	rawData := RawData{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&rawData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
	}
	defer r.Response.Body.Close()

	process(rawData)

	w.WriteHeader(http.StatusOK)
}
