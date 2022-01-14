package test

import (
	"context"
	"testing"

	"github.com/blockc0de/engine"
	"github.com/blockc0de/engine/interop"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestTimerNode(t *testing.T) {
	graphJson := `{"project_id":"51318fa9-656a-47ed-9c8c-1ced0a0793b9","name":"NEW_GRAPH","nodes":[{"id":"9c328401-c047-4c2e-ad0c-81b312c7c616","type":"*time.TimerNode","out_node":"18e96479-6ba9-471e-8404-d937ca1cc378","can_be_executed":false,"can_execute":true,"friendly_name":"Timer","block_type":"entry-point","_x":514,"_y":255,"in_parameters":[{"id":"3d97792e-ac3d-48ae-a869-031a9af84571","name":"intervalInSeconds","type":"System.Double","value":null,"assignment":"d17abe7e-79ff-4343-8c33-36f7b0b803bf","assignment_node":"d240e96a-401e-47ea-8108-c7987d3fa435","value_is_reference":false}],"out_parameters":[]},{"id":"d240e96a-401e-47ea-8108-c7987d3fa435","type":"*vars.DecimalNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Decimal","block_type":"variable","_x":252,"_y":336,"in_parameters":[],"out_parameters":[{"id":"d17abe7e-79ff-4343-8c33-36f7b0b803bf","name":"value","type":"System.Double","value":"3","assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"18e96479-6ba9-471e-8404-d937ca1cc378","type":"*console.PrintNode","out_node":null,"can_be_executed":true,"can_execute":true,"friendly_name":"Print","block_type":"function","_x":978,"_y":269,"in_parameters":[{"id":"fe315146-1b03-4bff-a865-cac3fef169e9","name":"message","type":"System.String","value":null,"assignment":"eb125c56-f74f-4493-9cf0-62218e0d85ac","assignment_node":"72d43947-b53d-42a7-9274-3d939123b5a5","value_is_reference":false}],"out_parameters":[]},{"id":"72d43947-b53d-42a7-9274-3d939123b5a5","type":"*vars.StringNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"String","block_type":"variable","_x":740,"_y":384,"in_parameters":[],"out_parameters":[{"id":"eb125c56-f74f-4493-9cf0-62218e0d85ac","name":"value","type":"System.String","value":"timer","assignment":"","assignment_node":"","value_is_reference":false}]}],"comments":[]}`

	graph, err := interop.LoadGraph([]byte(graphJson))
	assert.Nil(t, err)

	var count int
	var result string
	var e *engine.Engine
	event := engine.Event{
		AppendLog: func(msgType string, message string) {
			if result == "" {
				result = message
			}

			if message == "timer" {
				count++
				if count == 3 {
					e.Stop()
				}
			}
		},
	}
	e = engine.NewEngine(graph, common.Address{}, nil, event)
	err = e.Run(context.Background())
	assert.Nil(t, err)

	assert.Equal(t, result, "timer")
}
