package time

import (
	"context"
	"math/big"
	"reflect"
	"time"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	timerNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "TimerNode", FriendlyName: "Timer", NodeType: attributes.NodeTypeEnumEvent, GroupName: "Time", BlockLimitPerGraph: -1}}
	timerNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Start a timer that will init a new execution cycle, from in parameter specified time."}}
	timerNodeGasConfiguration = []interface{}{attributes.NodeGasConfiguration{BlockGasPrice: big.NewInt(10000000000000)}}
)

type TimerNode struct {
	block.NodeBase
	timer *time.Timer
}

func NewTimerNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(TimerNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())
	node.NodeData.IsEventNode = true

	intervalInSeconds, err := block.NewNodeParameter(node, "intervalInSeconds", block.NodeParameterTypeEnumDecimal, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(intervalInSeconds)

	return node, nil
}

func (n *TimerNode) CanExecute() bool {
	return true
}

func (n *TimerNode) SetupEvent(scheduler block.NodeScheduler) error {
	seconds, err := n.getIntervalInSeconds()
	if err != nil {
		return err
	}

	n.timer = time.NewTimer(seconds)

	go func() {
		for range n.timer.C {
			scheduler.AddCycle(n, nil)
		}
	}()

	scheduler.AddCycle(n, nil)
	return nil
}

func (n *TimerNode) BeginCycle(ctx context.Context, scheduler block.NodeScheduler) {
	scheduler.NextNode(ctx, n)

	if n.timer != nil {
		seconds, err := n.getIntervalInSeconds()
		if err != nil {
			scheduler.Stop()
		} else {
			n.timer.Reset(seconds)
		}
	}
}

func (n *TimerNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return timerNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return timerNodeGraphDescription
	case reflect.TypeOf(attributes.NodeGasConfiguration{}):
		return timerNodeGasConfiguration
	default:
		return nil
	}
}

func (n *TimerNode) OnExecution(context.Context, block.NodeScheduler) error {
	return nil
}

func (n *TimerNode) OnStop() error {
	if n.timer != nil {
		n.timer.Stop()
	}
	return nil
}

func (n *TimerNode) getIntervalInSeconds() (time.Duration, error) {
	var converter block.NodeParameterConverter
	intervalInSeconds, ok := converter.ToDecimal(n.Data().InParameters.Get("intervalInSeconds").ComputeValue())
	if !ok {
		return 0, block.ErrInvalidParameter{Name: "intervalInSeconds"}
	}

	seconds := intervalInSeconds.IntPart()
	if seconds < 0 {
		seconds = 1
	}
	return time.Second * time.Duration(seconds), nil
}
