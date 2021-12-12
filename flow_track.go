package flow_track

import (
	"fmt"
	"github.com/smart-think-app/flow-track/core"
	"github.com/smart-think-app/flow-track/status_enum"
	"time"
)

type iFlowBuilder interface {
	Init(flowId string, function string, source string, service string) iFlowBuilder
	SetDuration(startTime time.Time) iFlowBuilder
	SetAction(action string) iFlowBuilder
	SetMetadata(request interface{}, response interface{}) iFlowBuilder
	SetStatus(status_enum.Status) iFlowBuilder
	Send()
}
type flowBuilder struct {
	flowId          string
	function        string
	source          string
	service         string
	memoryAllocated string
	duration        float64
	status          string
	metadata        interface{}
	action          string
}

func (b *flowBuilder) Init(flowId string, function string, source string, service string) iFlowBuilder {

	b.flowId = flowId
	b.function = function
	b.source = source
	b.service = service

	return b
}

func (b *flowBuilder) SetMetadata(request interface{}, response interface{}) iFlowBuilder {

	b.metadata = map[string]interface{}{
		"Request":  request,
		"Response": response,
	}

	return b
}

func (b *flowBuilder) SetStatus(status status_enum.Status) iFlowBuilder {

	b.status = string(status)
	return b
}

func (b *flowBuilder) SetDuration(startTime time.Time) iFlowBuilder {
	now := time.Now()
	duration := now.Unix() - startTime.Unix()
	b.duration = float64(duration)/float64(1000)
	return b
}

func (b *flowBuilder) SetAction(action string) iFlowBuilder {
	b.action = action
	return b
}

func (b *flowBuilder) Send() {
	memoryUsage := core.GetMemUsage()
	b.memoryAllocated = fmt.Sprintf("%vMiB" , memoryUsage.MemoryAllocated)
}

func NewFlowTrack() *flowBuilder {
	return &flowBuilder{}
}