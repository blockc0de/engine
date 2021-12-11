package test

import (
	"context"
	"github.com/blockc0de/engine"
	"github.com/blockc0de/engine/loader"
	"github.com/stretchr/testify/assert"
	"math"
	"strconv"
	"testing"
	"time"
)

func TestTimeNodes(t *testing.T) {
	graphJson := `{"project_id":"51318fa9-656a-47ed-9c8c-1ced0a0793b9","name":"NEW_GRAPH","nodes":[{"id":"23eabf2c-cd70-4faa-9d37-38230ec14599","type":"*time.GetTimestampNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Get Timestamp","block_type":"function","_x":391,"_y":364,"in_parameters":[],"out_parameters":[{"id":"21260bbc-517e-4a67-bacf-2157eef7bd53","name":"timestamp","type":"System.Double","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"29a77438-db5e-4968-a94d-1ab6ca80b5d2","type":"*time.FormatTimestampNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Format Timestamp","block_type":"function","_x":773,"_y":395,"in_parameters":[{"id":"66f3f361-4a01-4272-9703-8d827909f8a1","name":"timestamp","type":"System.Double","value":null,"assignment":"21260bbc-517e-4a67-bacf-2157eef7bd53","assignment_node":"23eabf2c-cd70-4faa-9d37-38230ec14599","value_is_reference":false},{"id":"870d800c-e2c0-49d5-a66e-bbb6bb5db382","name":"format","type":"System.String","value":null,"assignment":"cc42a5ee-ac4b-4ca3-a2ce-ebd37248f9e0","assignment_node":"8f9d3ff9-a4b6-499b-ac5d-ee35d434f096","value_is_reference":false}],"out_parameters":[{"id":"7abb9a4d-6b10-4b18-994d-6f08a380dd7c","name":"dateString","type":"System.String","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"8f9d3ff9-a4b6-499b-ac5d-ee35d434f096","type":"*variable.StringNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"String","block_type":"variable","_x":548,"_y":511,"in_parameters":[],"out_parameters":[{"id":"cc42a5ee-ac4b-4ca3-a2ce-ebd37248f9e0","name":"value","type":"System.String","value":"%Y/%m/%d %H:%M:%S","assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"f907b229-df57-44ee-b8b7-de388e0fc078","type":"*time.ParseTimestampNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Parse Timestamp","block_type":"function","_x":1114,"_y":411,"in_parameters":[{"id":"2666b77e-d645-4545-995a-c65bd7b65440","name":"dateString","type":"System.String","value":null,"assignment":"7abb9a4d-6b10-4b18-994d-6f08a380dd7c","assignment_node":"29a77438-db5e-4968-a94d-1ab6ca80b5d2","value_is_reference":false},{"id":"b314189b-9dcf-4086-95ec-67dfd2b2c095","name":"format","type":"System.String","value":null,"assignment":"cc42a5ee-ac4b-4ca3-a2ce-ebd37248f9e0","assignment_node":"8f9d3ff9-a4b6-499b-ac5d-ee35d434f096","value_is_reference":false}],"out_parameters":[{"id":"b944ffc9-76c0-42c6-aefd-e2bed17735f8","name":"timestamp","type":"System.Double","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"cc70fb08-078b-440a-9874-cd7dbc1e4f4f","type":"*nodes.EntryPointNode","out_node":"c7941a5d-a4da-4903-bdc3-1b28adaa593b","can_be_executed":false,"can_execute":true,"friendly_name":"Entry Point","block_type":"entry-point","_x":1111,"_y":229,"in_parameters":[],"out_parameters":[]},{"id":"c7941a5d-a4da-4903-bdc3-1b28adaa593b","type":"*console.PrintNode","out_node":null,"can_be_executed":true,"can_execute":true,"friendly_name":"Print","block_type":"function","_x":1564,"_y":361,"in_parameters":[{"id":"f7912624-030d-4d96-bde6-f0e140329c9a","name":"message","type":"System.String","value":null,"assignment":"b944ffc9-76c0-42c6-aefd-e2bed17735f8","assignment_node":"f907b229-df57-44ee-b8b7-de388e0fc078","value_is_reference":false}],"out_parameters":[]}],"comments":[]}`

	graph, err := loader.LoadGraph([]byte(graphJson))
	assert.Nil(t, err)

	var result string
	var e *engine.Engine
	event := engine.Event{
		AppendLog: func(msgType string, message string) {
			if result == "" {
				e.Stop()
				result = message
			}
		},
	}
	e = engine.NewEngine(graph, event)
	e.Run(context.Background())

	ts, err := strconv.ParseInt(result, 10, 64)
	assert.Nil(t, err)
	assert.True(t, math.Abs(float64(ts-time.Now().Unix())) < 10)
}
