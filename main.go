package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/glebarez/go-sqlite"
)

var conn *sql.DB

func main() {
	var err error
	// 1. Define an http server
	// 2. Connect to a DB
	conn, err = sql.Open("sqlite", "./mgt.db")
	if err != nil {
		panic(err)
		return
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/devices", devicesHandler)
	http.ListenAndServe(":8080", nil)
}

func devicesHandler(w http.ResponseWriter, r *http.Request) {
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

type RegisterDeviceRequest struct {
	Name     string `json:name`
	DeviceId string `json:deviceId`
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

type DeviceDTO struct {
}

func getDevice(w http.ResponseWriter, r *http.Request) {

}

func getDevices(w http.ResponseWriter, r *http.Request) {
	var devices []DeviceDTO
	rows, _ := conn.Query("select * from devices")
	defer rows.Close()
	for rows.Next() {
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
