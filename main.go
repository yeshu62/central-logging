package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var logs = Logs{}

func logHandler(c *gin.Context) {
	switch c.Request.Method {
	case "POST":
		var log Log
		if err := c.ShouldBindJSON(&log); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id := logs.AddLog(log)
		c.JSON(http.StatusCreated, gin.H{"id": id})

	case "GET":
		var filter LogFilter
		serviceName := c.Query("serviceName")
		if serviceName != "" {
			filter.ServiceName = serviceName
		}
		logs := logs.QueryLogs(filter)
		c.JSON(http.StatusOK, logs)

	default:
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "method not allowed"})
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
		time.Sleep(time.Second)
	}
}

func main() {
	r := gin.Default()

	r.POST("/logs", logHandler)
	r.GET("/logs", logHandler)

	go service("Service1", "http://localhost:8081/logs")
	go service("Service2", "http://localhost:8081/logs")
	go service("Service3", "http://localhost:8081/logs")

	log.Fatal(r.Run(":8081"))
}
