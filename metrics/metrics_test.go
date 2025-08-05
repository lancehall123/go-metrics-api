package metrics

import (
	"testing"
)

func TestGetMetrics(t *testing.T) {
	m, err := GetMetrics()
	if err != nil {
		t.Fatalf("GetMetrics failed: %v", err)
	}

	if m.CPUPercent < 0 || m.CPUPercent > 100 {
		t.Errorf("Invalid CPU usage: %f", m.CPUPercent)
	}

	if m.MemoryUsed == 0 {
		t.Errorf("MemoryUsed should not be 0")
	}

	if m.MemoryTotal == 0 {
		t.Errorf("MemoryTotal should not be 0")
	}

	if m.MemoryUsage < 0 || m.MemoryUsage > 100 {
		t.Errorf("Invalid MemoryUsage: %f", m.MemoryUsage)
	}

	if m.Timestamp == "" {
		t.Errorf("Timestamp is empty")
	}
}
