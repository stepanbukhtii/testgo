package testgo

import (
	"fmt"
	"runtime"
	"time"
)

type MemTimeProfiler struct {
	m1, m2    runtime.MemStats
	startTime time.Time
	elapsed   time.Duration
}

func (c *MemTimeProfiler) Start() {
	runtime.ReadMemStats(&c.m1)
	c.startTime = time.Now()
}

func (c *MemTimeProfiler) Finish() {
	c.elapsed = time.Since(c.startTime)
	runtime.ReadMemStats(&c.m2)
}

func (c MemTimeProfiler) Size() string {
	totalSize := int64(c.m2.TotalAlloc - c.m1.TotalAlloc)

	if totalSize < 1000 {
		return fmt.Sprintf("%d B", totalSize)
	}
	div, exp := int64(1000), 0
	for n := totalSize / 1000; n >= 1000; n /= 1000 {
		div *= 1000
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(totalSize)/float64(div), "kMGTPE"[exp])
}

func (c MemTimeProfiler) Time() string {
	return c.elapsed.String()
}
