package main

import (
	"database/sql"
	"log"
	"metric-gathering-tool-service/api"
	"metric-gathering-tool-service/queue"
	"net/http"

	_ "github.com/glebarez/go-sqlite"
)

var conn *sql.DB
var pool = make(chan queue.Pool)

func main() {
	var err error
	conn, err = sql.Open("sqlite", "./mgt.db")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/devices", api.DevicesHandler)
	http.HandleFunc("/metric/send", api.MetricsHandelr)
	http.ListenAndServe(":8080", nil)
}
