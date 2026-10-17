package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DeviceDTO struct {
}

func getDevice(w http.ResponseWriter, r *http.Request) {

}

func getDevices(w http.ResponseWriter, r *http.Request) {
	var devices []DeviceDTO
	rows, _ := conn.Query("select * from devices")
	defer rows.Close()
	for rows.Next() {
		if rows.Err() != nil {
			return
		}
		var device DeviceDTO
		rows.Scan(&device)
		devices = append(devices, device)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(devices)
	if err != nil {
		return
	}
}

func deleteDevice(w http.ResponseWriter, r *http.Request) {

}

type RegisterDeviceRequest struct {
	Name     string `json:"name"`
	DeviceId string `json:"deviceId"`
}

func registerDevice(w http.ResponseWriter, r *http.Request) {
	var registerDeviceRequest RegisterDeviceRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&registerDeviceRequest)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer r.Body.Close()
	fmt.Println(registerDeviceRequest.Name, registerDeviceRequest.DeviceId)
	_, err = conn.Exec("insert into devices (name, device_id) values (?, ?)",
		registerDeviceRequest.Name,
		registerDeviceRequest.DeviceId,
	)
	if err != nil {
		return
	}
}
