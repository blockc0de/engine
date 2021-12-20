package test

import (
	"context"
	"testing"

	"github.com/blockc0de/engine"
	"github.com/blockc0de/engine/interop"
	"github.com/stretchr/testify/assert"
)

func TestVarsNode(t *testing.T) {
	graphJson := `{"project_id":"51318fa9-656a-47ed-9c8c-1ced0a0793b9","name":"NEW_GRAPH","nodes":[{"id":"9147506b-5d90-420a-8532-d6053f885b39","type":"*vars.StringNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"String","block_type":"variable","_x":332,"_y":394,"in_parameters":[],"out_parameters":[{"id":"19c3e211-c177-4781-8cb2-e73938d0b9c1","name":"value","type":"System.String","value":"a","assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"d8b03099-e486-4288-a6b4-6f6582f289c0","type":"*nodes.EntryPointNode","out_node":"ba382fa3-9afc-43ac-9f74-b03f63e63a24","can_be_executed":false,"can_execute":true,"friendly_name":"Entry Point","block_type":"entry-point","_x":156,"_y":288,"in_parameters":[],"out_parameters":[]},{"id":"8da6bc49-2cfe-471d-9b30-dcf7386e838c","type":"*vars.DecimalNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Decimal","block_type":"variable","_x":340,"_y":537,"in_parameters":[],"out_parameters":[{"id":"788714bb-7ac0-43ff-bfb3-a06325d0b8d8","name":"value","type":"System.Double","value":"10","assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"92813b7f-c32d-4e9a-ab46-cb36d72ae81b","type":"*vars.GetVariableNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Get variable","block_type":"function","_x":1190,"_y":212,"in_parameters":[{"id":"88f48378-056e-4cfc-9137-123e31206232","name":"name","type":"System.String","value":null,"assignment":"19c3e211-c177-4781-8cb2-e73938d0b9c1","assignment_node":"9147506b-5d90-420a-8532-d6053f885b39","value_is_reference":false}],"out_parameters":[{"id":"b987a14c-d860-4e96-854b-b14500a07103","name":"value","type":"System.Object","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"ba382fa3-9afc-43ac-9f74-b03f63e63a24","type":"*vars.SetVariableNode","out_node":"5131d734-ee75-4b80-a711-a0c983c1a686","can_be_executed":true,"can_execute":true,"friendly_name":"Set variable","block_type":"function","_x":817,"_y":403,"in_parameters":[{"id":"a22ace34-0ad3-4727-8e2c-3d2432379c89","name":"name","type":"System.String","value":null,"assignment":"19c3e211-c177-4781-8cb2-e73938d0b9c1","assignment_node":"9147506b-5d90-420a-8532-d6053f885b39","value_is_reference":false},{"id":"14c983e7-4282-40cd-936d-9cb110867cfd","name":"value","type":"System.Object","value":null,"assignment":"788714bb-7ac0-43ff-bfb3-a06325d0b8d8","assignment_node":"8da6bc49-2cfe-471d-9b30-dcf7386e838c","value_is_reference":false}],"out_parameters":[]},{"id":"8a16346c-8b75-457f-8a0c-c09c68df9200","type":"*math.AddNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Add A + B","block_type":"function","_x":1557,"_y":385,"in_parameters":[{"id":"18c7fa48-49d4-4bac-8209-80e5835980e4","name":"a","type":"System.Double","value":null,"assignment":"b987a14c-d860-4e96-854b-b14500a07103","assignment_node":"92813b7f-c32d-4e9a-ab46-cb36d72ae81b","value_is_reference":false},{"id":"182bba3e-5ab9-45f5-adb5-618ee61de79d","name":"b","type":"System.Double","value":null,"assignment":"6f222692-2d57-4f7c-a3e8-41332446a1ae","assignment_node":"d1641cfe-144f-4b5b-a3e8-a5a127d8e8b6","value_is_reference":false}],"out_parameters":[{"id":"6a5916de-ef00-4a6e-8b26-f3dddc0005b4","name":"value","type":"System.Double","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"d1641cfe-144f-4b5b-a3e8-a5a127d8e8b6","type":"*vars.DecimalNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Decimal","block_type":"variable","_x":1229,"_y":471,"in_parameters":[],"out_parameters":[{"id":"6f222692-2d57-4f7c-a3e8-41332446a1ae","name":"value","type":"System.Double","value":"20","assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"5131d734-ee75-4b80-a711-a0c983c1a686","type":"*vars.IsVariableExistNode","out_node":null,"can_be_executed":true,"can_execute":false,"friendly_name":"Is Variable Exist","block_type":"function","_x":1661,"_y":582,"in_parameters":[{"id":"4e7cb303-82e4-4572-9a04-fe6519cb2cbf","name":"name","type":"System.String","value":null,"assignment":"19c3e211-c177-4781-8cb2-e73938d0b9c1","assignment_node":"9147506b-5d90-420a-8532-d6053f885b39","value_is_reference":false}],"out_parameters":[{"id":"57f82eb1-0b2a-41c2-b61d-a020efd2d720","name":"true","type":"NodeBlock.Engine.Node","value":"f40bcf41-430c-4678-b411-4002a8f88abe","assignment":"","assignment_node":"","value_is_reference":true},{"id":"5ff9d5ab-4c6e-4dba-ba54-9899e0190d33","name":"false","type":"NodeBlock.Engine.Node","value":null,"assignment":"","assignment_node":"","value_is_reference":true}]},{"id":"f40bcf41-430c-4678-b411-4002a8f88abe","type":"*console.PrintNode","out_node":null,"can_be_executed":true,"can_execute":true,"friendly_name":"Print","block_type":"function","_x":2026,"_y":420,"in_parameters":[{"id":"b4e62d40-4bc1-4e74-9b7b-6f54dc4638e5","name":"message","type":"System.String","value":null,"assignment":"6a5916de-ef00-4a6e-8b26-f3dddc0005b4","assignment_node":"8a16346c-8b75-457f-8a0c-c09c68df9200","value_is_reference":false}],"out_parameters":[]}],"comments":[]}`

	graph, err := interop.LoadGraph([]byte(graphJson))
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
	err = e.Run(context.Background())
	assert.Nil(t, err)

	assert.Equal(t, result, "30")
}
