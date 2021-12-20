package test

import (
	"context"
	"testing"

	"github.com/blockc0de/engine"
	"github.com/blockc0de/engine/interop"
	"github.com/stretchr/testify/assert"
)

func TestArrayNode(t *testing.T) {
	graphJson := `{"project_id":"5469c406-bebb-44f9-ada4-2f4c08b00d98","name":"NEW_GRAPH","nodes":[{"id":"b7489b15-0f7c-43a6-80de-a238eb9b7513","type":"*array.CreateArrayNode","out_node":"9f36f737-470e-4123-94cd-c71f5ff3c588","can_be_executed":true,"can_execute":true,"friendly_name":"Create Array","block_type":"function","_x":526,"_y":441,"in_parameters":[],"out_parameters":[{"id":"687ff324-1e69-4b43-948e-b6c99f0cd2fd","name":"array","type":"System.Collections.Generic.List","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"50a40310-6917-4477-b696-c9ec8bc0145d","type":"*nodes.EntryPointNode","out_node":"b7489b15-0f7c-43a6-80de-a238eb9b7513","can_be_executed":false,"can_execute":true,"friendly_name":"Entry Point","block_type":"entry-point","_x":267,"_y":441,"in_parameters":[],"out_parameters":[]},{"id":"9f36f737-470e-4123-94cd-c71f5ff3c588","type":"*array.AddElementNode","out_node":"8b03b8fb-c090-4220-8f05-a5c98162d6a4","can_be_executed":true,"can_execute":true,"friendly_name":"Add Array Element","block_type":"function","_x":953,"_y":436,"in_parameters":[{"id":"7f202573-2ddf-40bc-a2ea-b7f5b66b906b","name":"array","type":"System.Collections.Generic.List","value":null,"assignment":"687ff324-1e69-4b43-948e-b6c99f0cd2fd","assignment_node":"b7489b15-0f7c-43a6-80de-a238eb9b7513","value_is_reference":false},{"id":"908002c3-1d11-4292-bcd3-2e446ec97ace","name":"element","type":"System.Object","value":null,"assignment":"b091ac15-8616-4406-a7fa-2201369cd30c","assignment_node":"d76ac42c-18a2-427d-8174-f57e6cda2db4","value_is_reference":false}],"out_parameters":[]},{"id":"d76ac42c-18a2-427d-8174-f57e6cda2db4","type":"*vars.StringNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"String","block_type":"variable","_x":684,"_y":570,"in_parameters":[],"out_parameters":[{"id":"b091ac15-8616-4406-a7fa-2201369cd30c","name":"value","type":"System.String","value":"1024","assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"8b03b8fb-c090-4220-8f05-a5c98162d6a4","type":"*array.ClearArrayNode","out_node":"1c1e7347-0da0-495a-bc45-8ca4d2951ee9","can_be_executed":true,"can_execute":true,"friendly_name":"Clear Array","block_type":"function","_x":1265,"_y":552,"in_parameters":[{"id":"be3f8d74-4ebf-452e-a867-3a69306c18e0","name":"array","type":"System.Collections.Generic.List","value":null,"assignment":"687ff324-1e69-4b43-948e-b6c99f0cd2fd","assignment_node":"b7489b15-0f7c-43a6-80de-a238eb9b7513","value_is_reference":false}],"out_parameters":[]},{"id":"1c1e7347-0da0-495a-bc45-8ca4d2951ee9","type":"*array.AddElementNode","out_node":"e815c31c-b431-4689-8136-0418393ccc27","can_be_executed":true,"can_execute":true,"friendly_name":"Add Array Element","block_type":"function","_x":1550,"_y":644,"in_parameters":[{"id":"670e5819-7b4b-4a62-b1a2-5e3e62646d59","name":"array","type":"System.Collections.Generic.List","value":null,"assignment":"687ff324-1e69-4b43-948e-b6c99f0cd2fd","assignment_node":"b7489b15-0f7c-43a6-80de-a238eb9b7513","value_is_reference":false},{"id":"7e70d2ec-5424-466e-8b55-f7491b31fa46","name":"element","type":"System.Object","value":null,"assignment":"3a8ada5c-fb7d-4553-b333-26800fc0364e","assignment_node":"5c746fdd-f5fd-47ab-9fe3-52737dc33b51","value_is_reference":false}],"out_parameters":[]},{"id":"5c746fdd-f5fd-47ab-9fe3-52737dc33b51","type":"*vars.StringNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"String","block_type":"variable","_x":1309,"_y":723,"in_parameters":[],"out_parameters":[{"id":"3a8ada5c-fb7d-4553-b333-26800fc0364e","name":"value","type":"System.String","value":"1","assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"e815c31c-b431-4689-8136-0418393ccc27","type":"*array.AddElementNode","out_node":"1e8afd31-a53e-45bd-8c12-ba90cee1dbe3","can_be_executed":true,"can_execute":true,"friendly_name":"Add Array Element","block_type":"function","_x":1837,"_y":646,"in_parameters":[{"id":"01744e8e-1ea5-4075-b0a9-643a2c487ad6","name":"array","type":"System.Collections.Generic.List","value":null,"assignment":"687ff324-1e69-4b43-948e-b6c99f0cd2fd","assignment_node":"b7489b15-0f7c-43a6-80de-a238eb9b7513","value_is_reference":false},{"id":"1e1f2866-5074-4553-92e4-c6067ebc1a29","name":"element","type":"System.Object","value":null,"assignment":"213f86a2-02b4-40d8-8eaa-4f81e3fa5e8b","assignment_node":"1c9ec055-9cea-4111-816e-9cee50a07650","value_is_reference":false}],"out_parameters":[]},{"id":"1c9ec055-9cea-4111-816e-9cee50a07650","type":"*vars.DecimalNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Decimal","block_type":"variable","_x":1596,"_y":798,"in_parameters":[],"out_parameters":[{"id":"213f86a2-02b4-40d8-8eaa-4f81e3fa5e8b","name":"value","type":"System.Double","value":"2","assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"57aa3be0-e0b4-4513-9205-733a7b4908d4","type":"*array.GetArraySizeNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Get Array Size","block_type":"function","_x":1980,"_y":783,"in_parameters":[{"id":"fcf4f415-dab8-4f00-86a0-aecd347c1019","name":"array","type":"System.Collections.Generic.List","value":null,"assignment":"687ff324-1e69-4b43-948e-b6c99f0cd2fd","assignment_node":"b7489b15-0f7c-43a6-80de-a238eb9b7513","value_is_reference":false}],"out_parameters":[{"id":"c47626a1-c153-4b8d-8201-6d643a73e170","name":"size","type":"System.Double","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"1e8afd31-a53e-45bd-8c12-ba90cee1dbe3","type":"*console.PrintNode","out_node":"3602d434-7c6a-4fbf-962e-91f753f09d15","can_be_executed":true,"can_execute":true,"friendly_name":"Print","block_type":"function","_x":2193,"_y":659,"in_parameters":[{"id":"4b7b235e-0357-4e1c-8217-a584d538c28b","name":"message","type":"System.String","value":null,"assignment":"c47626a1-c153-4b8d-8201-6d643a73e170","assignment_node":"57aa3be0-e0b4-4513-9205-733a7b4908d4","value_is_reference":false}],"out_parameters":[]},{"id":"fdbe7e48-f071-4754-8f3e-a83bb17ba414","type":"*array.GetElementAtIndexNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Get Array Element At Index","block_type":"function","_x":2477,"_y":796,"in_parameters":[{"id":"0af28907-50fa-4891-8eab-c9949bb5b48e","name":"array","type":"System.Collections.Generic.List","value":null,"assignment":"687ff324-1e69-4b43-948e-b6c99f0cd2fd","assignment_node":"b7489b15-0f7c-43a6-80de-a238eb9b7513","value_is_reference":false},{"id":"0943ac4f-92ca-4faf-963b-a763fe17132d","name":"index","type":"System.Double","value":null,"assignment":"8b3a067c-1431-424d-ac68-446125bb514a","assignment_node":"982cf61a-22c9-41f9-a50d-8e498cb690fc","value_is_reference":false}],"out_parameters":[{"id":"03b8031b-7852-419d-90be-aff0dbe4e219","name":"element","type":"System.Object","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"982cf61a-22c9-41f9-a50d-8e498cb690fc","type":"*vars.DecimalNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Decimal","block_type":"variable","_x":2200,"_y":891,"in_parameters":[],"out_parameters":[{"id":"8b3a067c-1431-424d-ac68-446125bb514a","name":"value","type":"System.Double","value":"1","assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"3602d434-7c6a-4fbf-962e-91f753f09d15","type":"*console.PrintNode","out_node":"b58830d6-c971-42d9-a24e-f960a199b301","can_be_executed":true,"can_execute":true,"friendly_name":"Print","block_type":"function","_x":2743,"_y":694,"in_parameters":[{"id":"b8401a66-d03f-460e-9022-27c1cd3c3217","name":"message","type":"System.String","value":null,"assignment":"03b8031b-7852-419d-90be-aff0dbe4e219","assignment_node":"fdbe7e48-f071-4754-8f3e-a83bb17ba414","value_is_reference":false}],"out_parameters":[]},{"id":"b58830d6-c971-42d9-a24e-f960a199b301","type":"*array.EachElementArrayNode","out_node":"24aae804-5b19-4c83-8e59-f7168c6b2e53","can_be_executed":true,"can_execute":true,"friendly_name":"Each Element In Array","block_type":"function","_x":3031,"_y":766,"in_parameters":[{"id":"3de96483-d835-45a7-8ae8-3d5fa6d5f747","name":"array","type":"System.Collections.Generic.List","value":null,"assignment":"687ff324-1e69-4b43-948e-b6c99f0cd2fd","assignment_node":"b7489b15-0f7c-43a6-80de-a238eb9b7513","value_is_reference":false}],"out_parameters":[{"id":"2e577478-a322-4e10-8e57-d89437ab32a3","name":"each","type":"NodeBlock.Engine.Node","value":"f1d931b9-f981-4a8e-8a44-5898967a060c","assignment":"","assignment_node":"","value_is_reference":true},{"id":"f78102a1-6588-4044-80b7-e7aa1ce3d29a","name":"item","type":"System.Object","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"f1d931b9-f981-4a8e-8a44-5898967a060c","type":"*console.PrintNode","out_node":null,"can_be_executed":true,"can_execute":true,"friendly_name":"Print","block_type":"function","_x":3340,"_y":896,"in_parameters":[{"id":"f35dbe2b-15cb-4c79-9396-5ecfc1f16859","name":"message","type":"System.String","value":null,"assignment":"f78102a1-6588-4044-80b7-e7aa1ce3d29a","assignment_node":"b58830d6-c971-42d9-a24e-f960a199b301","value_is_reference":false}],"out_parameters":[]},{"id":"665d4741-9587-47a5-9fd7-024c629a33a1","type":"*encoding.ConvertToJsonNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Convert To JSON","block_type":"function","_x":3592,"_y":817,"in_parameters":[{"id":"8760ac1f-9dce-4c71-b6af-5d868e1410a3","name":"object","type":"System.Object","value":null,"assignment":"687ff324-1e69-4b43-948e-b6c99f0cd2fd","assignment_node":"b7489b15-0f7c-43a6-80de-a238eb9b7513","value_is_reference":false}],"out_parameters":[{"id":"316fe51c-3e26-4435-920e-25cfb835866f","name":"json","type":"System.String","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"24aae804-5b19-4c83-8e59-f7168c6b2e53","type":"*console.PrintNode","out_node":null,"can_be_executed":true,"can_execute":true,"friendly_name":"Print","block_type":"function","_x":3824,"_y":735,"in_parameters":[{"id":"5ffa2a4d-342d-4e99-91b5-ae0b4416447d","name":"message","type":"System.String","value":null,"assignment":"316fe51c-3e26-4435-920e-25cfb835866f","assignment_node":"665d4741-9587-47a5-9fd7-024c629a33a1","value_is_reference":false}],"out_parameters":[]}],"comments":[]}`

	graph, err := interop.LoadGraph([]byte(graphJson))
	assert.Nil(t, err)

	var e *engine.Engine
	result := make([]string, 0)
	event := engine.Event{
		AppendLog: func(msgType string, message string) {
			if len(result) < 5 {
				result = append(result, message)
				if len(result) == 5 {
					e.Stop()
				}
			}
		},
	}
	e = engine.NewEngine(graph, event)
	err = e.Run(context.Background())
	assert.Nil(t, err)

	assert.Equal(t, result, []string{"2", "2", "1", "2", "[\"1\",\"2\"]"})
}
