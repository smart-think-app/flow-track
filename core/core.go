package core

import (
	"github.com/smart-think-app/flow-track/dto"
	"runtime"
)

func GetMemUsage() dto.MemoryUsageDto {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return dto.MemoryUsageDto{MemoryAllocated: bToMb(m.Sys)}
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}