package flow_track

import (
	"github.com/smart-think-app/flow-track/core"
	"time"
)

func Track(f func() error) {
	startTime := time.Now()
	if err := f(); err != nil {

	} else {
		core.PrintMemUsage()
		core.PrintDuration(startTime)
	}
}
