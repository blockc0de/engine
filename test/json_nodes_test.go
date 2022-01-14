package test

import (
	"context"
	"testing"

	"github.com/blockc0de/engine"
	"github.com/blockc0de/engine/interop"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestJsonNode(t *testing.T) {
	graphJson := `{"project_id":"95b7f01c-9b25-4eee-b39c-f236b45652f3","name":"NEW_GRAPH","nodes":[{"id":"707c6a42-cfdc-45dc-b775-f5300936cf1b","type":"*nodes.EntryPointNode","out_node":"3d97aa15-91a8-49e4-849a-c282fd238489","can_be_executed":false,"can_execute":true,"friendly_name":"Entry Point","block_type":"entry-point","_x":284,"_y":322,"in_parameters":[],"out_parameters":[]},{"id":"3d97aa15-91a8-49e4-849a-c282fd238489","type":"*encoding.CreateJsonObjectNode","out_node":"4796aa97-6ebe-48db-8686-5d6b72e8a267","can_be_executed":true,"can_execute":true,"friendly_name":"Create JSON Object","block_type":"function","_x":572,"_y":353,"in_parameters":[],"out_parameters":[{"id":"0905b441-226e-4496-8918-0b7584acfa67","name":"jsonObject","type":"System.Object","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"4796aa97-6ebe-48db-8686-5d6b72e8a267","type":"*encoding.AddJsonValueNode","out_node":"9938f680-b057-48ba-8d1c-f9b07deddc55","can_be_executed":true,"can_execute":true,"friendly_name":"Add JSON Property","block_type":"function","_x":941,"_y":368,"in_parameters":[{"id":"76cc9842-adb7-42e6-97f8-d924dcd6a385","name":"jsonObject","type":"System.Object","value":null,"assignment":"0905b441-226e-4496-8918-0b7584acfa67","assignment_node":"3d97aa15-91a8-49e4-849a-c282fd238489","value_is_reference":false},{"id":"07784a06-c2ce-44b2-95a8-fc3c9531b5a4","name":"key","type":"System.String","value":null,"assignment":"02d1f33a-d12a-4900-863b-a13868c65d35","assignment_node":"a45408be-f6f1-4ff9-b614-c91be41903b4","value_is_reference":false},{"id":"379f44f7-a169-4b2e-8405-fb60e5e7d8e2","name":"value","type":"System.String","value":null,"assignment":"1227d41b-963b-4a11-991f-2519b24eb637","assignment_node":"e7e3c903-bbb8-4b61-ae87-e7542b4b83c7","value_is_reference":false}],"out_parameters":[{"id":"2f5f1ef8-1999-474b-ae8c-16e03bd90116","name":"jsonObjectOut","type":"System.Object","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"a45408be-f6f1-4ff9-b614-c91be41903b4","type":"*vars.StringNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"String","block_type":"variable","_x":594,"_y":472,"in_parameters":[],"out_parameters":[{"id":"02d1f33a-d12a-4900-863b-a13868c65d35","name":"value","type":"System.String","value":"key","assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"e7e3c903-bbb8-4b61-ae87-e7542b4b83c7","type":"*vars.BoolNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Boolean","block_type":"variable","_x":617,"_y":641,"in_parameters":[],"out_parameters":[{"id":"1227d41b-963b-4a11-991f-2519b24eb637","name":"value","type":"System.Boolean","value":"1","assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"502ccd6a-2e32-4a08-bd6d-91a6d57f1a64","type":"*encoding.ConvertToJsonNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Convert To JSON","block_type":"function","_x":1263,"_y":502,"in_parameters":[{"id":"0621835b-198b-49cf-8303-9a824849ad2c","name":"object","type":"System.Object","value":null,"assignment":"2f5f1ef8-1999-474b-ae8c-16e03bd90116","assignment_node":"4796aa97-6ebe-48db-8686-5d6b72e8a267","value_is_reference":false}],"out_parameters":[{"id":"97fd02db-c331-4151-84b7-6d3daa812bad","name":"json","type":"System.String","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"c73b3117-f550-4c82-ba51-6fb00ec1cbd1","type":"*encoding.JsonToJsonObjectNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"JSON to JSON Object","block_type":"function","_x":1515,"_y":424,"in_parameters":[{"id":"a43929b2-addd-4f40-889f-2e09c44a2cc6","name":"json","type":"System.String","value":null,"assignment":"97fd02db-c331-4151-84b7-6d3daa812bad","assignment_node":"502ccd6a-2e32-4a08-bd6d-91a6d57f1a64","value_is_reference":false}],"out_parameters":[{"id":"0f4a25ee-bf91-465b-b1d6-53ded1181572","name":"jsonObject","type":"System.Object","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"e5463edb-683c-47b7-a179-aabe03274df8","type":"*encoding.ConvertToJsonNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Convert To JSON","block_type":"function","_x":1768,"_y":412,"in_parameters":[{"id":"394fd975-880f-4ec7-b4c0-dad41cd356b8","name":"object","type":"System.Object","value":null,"assignment":"0f4a25ee-bf91-465b-b1d6-53ded1181572","assignment_node":"c73b3117-f550-4c82-ba51-6fb00ec1cbd1","value_is_reference":false}],"out_parameters":[{"id":"f340a8f0-b0dd-40b6-9414-28276e247e60","name":"json","type":"System.String","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"9938f680-b057-48ba-8d1c-f9b07deddc55","type":"*encoding.MergeJsonNodeNode","out_node":"dee3445a-8781-4f12-9c3c-af241371ea8d","can_be_executed":true,"can_execute":true,"friendly_name":"Merge JSON","block_type":"function","_x":2113,"_y":361,"in_parameters":[{"id":"faaf0344-55c2-43ee-a93f-9413887f0a60","name":"json1","type":"System.String","value":null,"assignment":"f340a8f0-b0dd-40b6-9414-28276e247e60","assignment_node":"e5463edb-683c-47b7-a179-aabe03274df8","value_is_reference":false},{"id":"d24b9b75-42db-41e6-b9bb-fefa91fcf530","name":"json2","type":"System.String","value":null,"assignment":"ccf2481b-e9ec-437d-958c-ec604f9bf812","assignment_node":"b4699351-8982-449e-aee5-8d4ade8fcbc2","value_is_reference":false}],"out_parameters":[{"id":"63c83428-7277-4f59-a0f2-34875c63498d","name":"mergedJson","type":"System.String","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"b4699351-8982-449e-aee5-8d4ade8fcbc2","type":"*vars.StringNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"String","block_type":"variable","_x":1845,"_y":510,"in_parameters":[],"out_parameters":[{"id":"ccf2481b-e9ec-437d-958c-ec604f9bf812","name":"value","type":"System.String","value":"{\"hello\":\"1234\"}","assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"dee3445a-8781-4f12-9c3c-af241371ea8d","type":"*encoding.JsonSelectorNodeNode","out_node":"e102ca7c-cf76-46a2-a080-dc62d32d44d2","can_be_executed":true,"can_execute":true,"friendly_name":"JSON Selector","block_type":"function","_x":2463,"_y":364,"in_parameters":[{"id":"5fdb71c0-4d6d-472a-8ccc-e8120b896f4b","name":"json","type":"System.String","value":null,"assignment":"63c83428-7277-4f59-a0f2-34875c63498d","assignment_node":"9938f680-b057-48ba-8d1c-f9b07deddc55","value_is_reference":false},{"id":"7d21ecd1-ee95-48c4-9cd4-721310815544","name":"selector","type":"System.String","value":null,"assignment":"7b3835ec-4623-4c9e-8532-42f20002694e","assignment_node":"fd012995-8d97-444a-b7f1-5c3928923ea4","value_is_reference":false}],"out_parameters":[{"id":"dd82e9e1-4959-4438-a203-6fec4c8f16dc","name":"value","type":"System.String","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"fd012995-8d97-444a-b7f1-5c3928923ea4","type":"*vars.StringNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"String","block_type":"variable","_x":2226,"_y":514,"in_parameters":[],"out_parameters":[{"id":"7b3835ec-4623-4c9e-8532-42f20002694e","name":"value","type":"System.String","value":"hello","assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"e102ca7c-cf76-46a2-a080-dc62d32d44d2","type":"*console.PrintNode","out_node":null,"can_be_executed":true,"can_execute":true,"friendly_name":"Print","block_type":"function","_x":2858,"_y":376,"in_parameters":[{"id":"a391280d-8e17-4702-8892-ffe99a7efdf7","name":"message","type":"System.String","value":null,"assignment":"dd82e9e1-4959-4438-a203-6fec4c8f16dc","assignment_node":"dee3445a-8781-4f12-9c3c-af241371ea8d","value_is_reference":false}],"out_parameters":[]}],"comments":[]}`

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
	e = engine.NewEngine(graph, common.Address{}, nil, event)
	err = e.Run(context.Background())
	assert.Nil(t, err)

	assert.Equal(t, result, "1234")
}
