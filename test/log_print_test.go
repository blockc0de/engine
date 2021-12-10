package test

import (
	"context"
	"github.com/blockc0de/engine"
	"github.com/blockc0de/engine/loader"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogPrint(t *testing.T) {
	graphJson := `{"project_id":"83315aa0-3733-4b49-8355-ff2785a26e7f","name":"NEW_GRAPH","nodes":[{"id":"ffca7476-9e85-4d3d-9713-6af2b4ccdcd4","type":"*nodes.EntryPointNode","out_node":"877ae8a0-5ab5-4ba1-bcc3-acd9060cd981","can_be_executed":false,"can_execute":true,"friendly_name":"Entry Point","block_type":"entry-point","_x":274,"_y":267,"in_parameters":[],"out_parameters":[]},{"id":"877ae8a0-5ab5-4ba1-bcc3-acd9060cd981","type":"*console.PrintNode","out_node":null,"can_be_executed":true,"can_execute":true,"friendly_name":"Print","block_type":"function","_x":730,"_y":290,"in_parameters":[{"id":"d0f64d68-6773-4795-9b79-347abcbc293c","name":"message","type":"System.String","value":null,"assignment":"d9568289-b4b2-40ec-bba0-f01b07064d80","assignment_node":"bd449051-e219-4cbf-b36e-ea1ffb9b5360","value_is_reference":false}],"out_parameters":[]},{"id":"bd449051-e219-4cbf-b36e-ea1ffb9b5360","type":"*variable.StringNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"String","block_type":"variable","_x":437,"_y":420,"in_parameters":[],"out_parameters":[{"id":"d9568289-b4b2-40ec-bba0-f01b07064d80","name":"value","type":"System.String","value":"hello world!","assignment":"","assignment_node":"","value_is_reference":false}]}],"comments":[]}`

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

	assert.Equal(t, result, "hello world!")
}
