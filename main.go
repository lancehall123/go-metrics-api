package main

import (
	"encoding/json"
	"log"
	"net/http"

	"go-metrics-api/metrics"
)

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	m, err := metrics.GetMetrics()
	if err != nil {
		http.Error(w, "Failed to get metrics", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(m)
}

func main() {
	http.HandleFunc("/metrics", metricsHandler)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
