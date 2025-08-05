package metrics

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type Metrics struct {
	CPUPercent  float64 `json:"cpu_percent"`
	MemoryUsed  uint64  `json:"memory_used_mb"`
	MemoryTotal uint64  `json:"memory_total_mb"`
	MemoryUsage float64 `json:"memory_usage_percent"`
	Timestamp   string  `json:"timestamp"`
}

func GetMetrics() (*Metrics, error) {
	cpuPercents, err := cpu.Percent(0, false)
	if err != nil {
		return nil, err
	}
	memStats, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	return &Metrics{
		CPUPercent:  cpuPercents[0],
		MemoryUsed:  memStats.Used / 1024 / 1024,
		MemoryTotal: memStats.Total / 1024 / 1024,
		MemoryUsage: memStats.UsedPercent,
		Timestamp:   time.Now().Format(time.RFC3339),
	}, nil
}
