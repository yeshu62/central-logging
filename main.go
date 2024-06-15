package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var logs = Logs{}

func logHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var log Log
		if err := json.NewDecoder(r.Body).Decode(&log); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id := logs.AddLog(log)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]int64{"id": id})

	case "GET":
		var filter LogFilter
		if serviceName := r.URL.Query().Get("serviceName"); serviceName != "" {
			filter.ServiceName = serviceName
		}
		logs := logs.QueryLogs(filter)
		json.NewEncoder(w).Encode(logs)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func service(serviceName string, url string) {
	fmt.Println("Post req from service: ", serviceName)

	for {
		log := Log{
			ServiceName: serviceName,
			Severity:    []string{"INFO", "WARN", "ERROR"}[rand.Intn(3)],
			Message:     "...log message",
		}
		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(log)
		_, err := http.Post(url, "application/json", buf)
		if err != nil {
			fmt.Println("Error logging", err)
		}
		time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	}
}

func main() {
	go service("Service1", "http://localhost:8080/logs")
	go service("Service2", "http://localhost:8080/logs")
	go service("Service3", "http://localhost:8080/logs")

	http.HandleFunc("/logs", logHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
