package api

import (
	"net/http"
	"strings"
)

func DevicesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if strings.Contains(r.URL.Path, "devices") {
			getDevices(w, r)
		} else {
			getDevice(w, r)
		}
	} else if r.Method == "POST" {
		registerDevice(w, r)
	} else if r.Method == "DELETE" {
		deleteDevice(w, r)
	}
}

func MetricsHandelr(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if strings.Contains(r.URL.Path, "send") {
			send(w, r)
		}
	}
}
