package main

import (
	"sync"
	"time"
)

type Log struct {
	ID          int64     `json:"id"`
	Timestamp   time.Time `json:"timestamp"`
	Severity    string    `json:"severity"`
	ServiceName string    `json:"serviceName"`
	Message     string    `json:"message"`
}

type Logs struct {
	logs   []Log
	mu     sync.Mutex
	nextId int64
}

type LogFilter struct {
	Timestamp   time.Time
	Severity    string
	ServiceName string
}

func (logs *Logs) AddLog(log Log) int64 {
	logs.mu.Lock()
	defer logs.mu.Unlock()
	log.ID = logs.nextId
	logs.nextId++
	log.Timestamp = time.Now()
	logs.logs = append(logs.logs, log)
	return log.ID
}

func (f LogFilter) Matches(log Log) bool {
	return (f.Timestamp.IsZero() || log.Timestamp.After(f.Timestamp)) && (f.Severity == "" || f.Severity == log.Severity) && (f.ServiceName == "" || f.ServiceName == log.ServiceName)
}

func (logs *Logs) QueryLogs(filter LogFilter) []Log {
	logs.mu.Lock()
	defer logs.mu.Unlock()
	var result []Log
	for _, log := range logs.logs {
		if filter.Matches(log) {
			result = append(result, log)
		}
	}
	return result
}
