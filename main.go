package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type Metrics struct {
	CPUPercent float64 `json:"cpu_percent"`
	MemoryUsed uint64  `json:"memory_used_mb"`
	MemoryTotal uint64 `json:"memory_total_mb"`
	MemoryUsage float64 `json:"memory_usage_percent"`
	Timestamp   string  `json:"timestamp"`
}

func getMetrics() (*Metrics, error) {
	cpuPercents, err := cpu.Percent(0, false)
	if err != nil {
		return nil, err
	}
	memStats, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	return &Metrics{
		CPUPercent: cpuPercents[0],
		MemoryUsed: memStats.Used / 1024 / 1024,
		MemoryTotal: memStats.Total / 1024 / 1024,
		MemoryUsage: memStats.UsedPercent,
		Timestamp:   time.Now().Format(time.RFC3339),
	}, nil
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	metrics, err := getMetrics()
	if err != nil {
		http.Error(w, "Failed to get metrics", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}

func main() {
	http.HandleFunc("/metrics", metricsHandler)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
