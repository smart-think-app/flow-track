package core

import (
	"fmt"
	"runtime"
	"time"
)

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v", m.NumGC)
}

func PrintDuration(startTime time.Time) {
	now := time.Now()
	duration := now.Unix() - startTime.Unix()
	fmt.Printf("\tDuration = %d\n" ,duration )
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}