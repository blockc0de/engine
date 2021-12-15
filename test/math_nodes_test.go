package test

import (
	"context"
	"testing"

	"github.com/blockc0de/engine"
	"github.com/blockc0de/engine/interop"
	"github.com/stretchr/testify/assert"
)

func TestMathNodes(t *testing.T) {
	graphJson := `{"project_id":"488ebc23-4d43-4db8-a025-781a97224bd9","name":"NEW_GRAPH","nodes":[{"id":"4a741489-fe4f-4051-9bac-38ff28b2189b","type":"*math.AddNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Add A + B","block_type":"function","_x":536,"_y":360,"in_parameters":[{"id":"b927d938-033c-487f-b5cb-300c8861c330","name":"a","type":"System.Double","value":null,"assignment":"5f5a94d0-4303-471b-88c4-3eda7b701f60","assignment_node":"dcd77db5-6e1d-49db-8662-82ef9de76157","value_is_reference":false},{"id":"81d9da4d-a5c7-466b-a70c-8cd4b1797592","name":"b","type":"System.Double","value":null,"assignment":"3c55be52-1c45-4f25-85f7-d3ebfda792ee","assignment_node":"4e13060f-6ed9-4493-b96b-55bca3c4eb7a","value_is_reference":false}],"out_parameters":[{"id":"46cf2cf7-98f3-41b3-8c04-f155311ff3ec","name":"value","type":"System.Double","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"dcd77db5-6e1d-49db-8662-82ef9de76157","type":"*vars.DecimalNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Decimal","block_type":"variable","_x":212,"_y":268,"in_parameters":[],"out_parameters":[{"id":"5f5a94d0-4303-471b-88c4-3eda7b701f60","name":"value","type":"System.Double","value":"1.2","assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"4e13060f-6ed9-4493-b96b-55bca3c4eb7a","type":"*vars.DecimalNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Decimal","block_type":"variable","_x":211,"_y":423,"in_parameters":[],"out_parameters":[{"id":"3c55be52-1c45-4f25-85f7-d3ebfda792ee","name":"value","type":"System.Double","value":"2.3","assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"1cb57926-dda9-4e40-b5a0-2994683f791d","type":"*math.SubNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Substract A - B","block_type":"function","_x":978,"_y":376,"in_parameters":[{"id":"357843a1-6f78-48c3-a837-b3d285da12f8","name":"a","type":"System.Double","value":null,"assignment":"46cf2cf7-98f3-41b3-8c04-f155311ff3ec","assignment_node":"4a741489-fe4f-4051-9bac-38ff28b2189b","value_is_reference":false},{"id":"4091bb89-de30-4504-bff1-1d351321c3b1","name":"b","type":"System.Double","value":null,"assignment":"4cc32b4e-5df9-49a7-8dda-c0d956215859","assignment_node":"45713222-70ef-4898-b148-151001a212e3","value_is_reference":false}],"out_parameters":[{"id":"bcfcf3a1-c332-4fe0-8710-203637f09306","name":"value","type":"System.Double","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"45713222-70ef-4898-b148-151001a212e3","type":"*vars.DecimalNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Decimal","block_type":"variable","_x":703,"_y":476,"in_parameters":[],"out_parameters":[{"id":"4cc32b4e-5df9-49a7-8dda-c0d956215859","name":"value","type":"System.Double","value":"0.1","assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"fc7a8b45-60ae-4e1c-b674-d0c0bc740983","type":"*math.MulNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Multiply A * B","block_type":"function","_x":1355,"_y":375,"in_parameters":[{"id":"d5f158e0-046d-4f4b-97a0-dddcf4bd8d3f","name":"a","type":"System.Double","value":null,"assignment":"bcfcf3a1-c332-4fe0-8710-203637f09306","assignment_node":"1cb57926-dda9-4e40-b5a0-2994683f791d","value_is_reference":false},{"id":"16d4796b-433c-47f1-b9e8-cb20d030ab95","name":"b","type":"System.Double","value":null,"assignment":"f408fc66-6d56-4e7b-b370-207e80d9b88c","assignment_node":"f730d4b7-f112-4ab8-817f-bfebc7fc7913","value_is_reference":false}],"out_parameters":[{"id":"3632b70e-5a82-465e-9327-6691888a81d6","name":"value","type":"System.Double","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"f730d4b7-f112-4ab8-817f-bfebc7fc7913","type":"*vars.DecimalNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Decimal","block_type":"variable","_x":1096,"_y":520,"in_parameters":[],"out_parameters":[{"id":"f408fc66-6d56-4e7b-b370-207e80d9b88c","name":"value","type":"System.Double","value":"2","assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"30e688ff-4a51-4382-b4fd-352c92224bb4","type":"*math.DivNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Divide A / B","block_type":"function","_x":1719,"_y":389,"in_parameters":[{"id":"43af4f5b-581c-495f-83db-b7261df2abbd","name":"a","type":"System.Double","value":null,"assignment":"3632b70e-5a82-465e-9327-6691888a81d6","assignment_node":"fc7a8b45-60ae-4e1c-b674-d0c0bc740983","value_is_reference":false},{"id":"c03b72b8-26ee-44e4-a13b-4be8fb7ecf50","name":"b","type":"System.Double","value":null,"assignment":"57b2df1d-1939-432e-b616-ceda2a5ba566","assignment_node":"98ab47ea-c937-4f89-8b55-f9b196204c42","value_is_reference":false}],"out_parameters":[{"id":"732aad10-f5a2-4bba-a0f5-8f9f7dcb3005","name":"value","type":"System.Double","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"98ab47ea-c937-4f89-8b55-f9b196204c42","type":"*vars.DecimalNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Decimal","block_type":"variable","_x":1462,"_y":525,"in_parameters":[],"out_parameters":[{"id":"57b2df1d-1939-432e-b616-ceda2a5ba566","name":"value","type":"System.Double","value":"3.4","assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"8a1aef70-dd82-479e-a561-90809dbbba9f","type":"*math.PowNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Power A ^ B","block_type":"function","_x":2023,"_y":424,"in_parameters":[{"id":"74724643-49c3-4d4c-8b66-ceea79411aea","name":"a","type":"System.Double","value":null,"assignment":"732aad10-f5a2-4bba-a0f5-8f9f7dcb3005","assignment_node":"30e688ff-4a51-4382-b4fd-352c92224bb4","value_is_reference":false},{"id":"926e1d93-516c-449b-a97f-8c300b82da98","name":"b","type":"System.Double","value":null,"assignment":"3f911d57-d430-4c27-a11d-b4f2710199f3","assignment_node":"507f5f9d-219a-471d-90e1-ca5d0bdd88db","value_is_reference":false}],"out_parameters":[{"id":"0cb5a3b0-52b8-4c47-b5c4-810c8ec5555f","name":"value","type":"System.Double","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"507f5f9d-219a-471d-90e1-ca5d0bdd88db","type":"*vars.DecimalNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Decimal","block_type":"variable","_x":1787,"_y":557,"in_parameters":[],"out_parameters":[{"id":"3f911d57-d430-4c27-a11d-b4f2710199f3","name":"value","type":"System.Double","value":"8","assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"a2034c95-8284-465b-8262-aafc5b6a0435","type":"*math.ModNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Modulo A % B","block_type":"function","_x":2344,"_y":460,"in_parameters":[{"id":"57703ce0-965f-4dff-b501-d0ba556c8bc8","name":"a","type":"System.Double","value":null,"assignment":"0cb5a3b0-52b8-4c47-b5c4-810c8ec5555f","assignment_node":"8a1aef70-dd82-479e-a561-90809dbbba9f","value_is_reference":false},{"id":"7be3e291-826b-46ec-9d46-d4eaf2369f74","name":"b","type":"System.Double","value":null,"assignment":"bcd1d633-8feb-44ad-acc4-08372f9cd0a1","assignment_node":"bdbb9d79-752f-4009-8238-396b0dd4253a","value_is_reference":false}],"out_parameters":[{"id":"90411389-57d9-45ee-b12d-20545ccdada1","name":"value","type":"System.Double","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"bdbb9d79-752f-4009-8238-396b0dd4253a","type":"*vars.DecimalNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Decimal","block_type":"variable","_x":2099,"_y":589,"in_parameters":[],"out_parameters":[{"id":"bcd1d633-8feb-44ad-acc4-08372f9cd0a1","name":"value","type":"System.Double","value":"200","assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"9a83df14-bb8e-48e7-9c55-d70cc91653a0","type":"*nodes.EntryPointNode","out_node":"339edee5-144c-4069-a4f4-2202350e3a67","can_be_executed":false,"can_execute":true,"friendly_name":"Entry Point","block_type":"entry-point","_x":226,"_y":66,"in_parameters":[],"out_parameters":[]},{"id":"6ad3c998-1346-4643-ba70-5bda986661e3","type":"*math.AddNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Add A + B","block_type":"function","_x":2747,"_y":484,"in_parameters":[{"id":"384e82b6-28f3-4e8b-8000-ac9ae1e983e5","name":"a","type":"System.Double","value":null,"assignment":"90411389-57d9-45ee-b12d-20545ccdada1","assignment_node":"a2034c95-8284-465b-8262-aafc5b6a0435","value_is_reference":false},{"id":"afbd2c16-6dfd-4771-ac97-54972cb31ecb","name":"b","type":"System.Double","value":null,"assignment":"ba22f9a9-00e1-44a8-81b4-11a3cc72b2dd","assignment_node":"2735022b-0e7b-4a25-a5db-7e028792043b","value_is_reference":false}],"out_parameters":[{"id":"58cc0ac5-3ae1-4b56-b85f-270ea565fa07","name":"value","type":"System.Double","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"2735022b-0e7b-4a25-a5db-7e028792043b","type":"*vars.DecimalNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Decimal","block_type":"variable","_x":2471,"_y":613,"in_parameters":[],"out_parameters":[{"id":"ba22f9a9-00e1-44a8-81b4-11a3cc72b2dd","name":"value","type":"System.Double","value":"1.234","assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"9d014de1-d97e-41ff-af62-defc2ca3bcf1","type":"*math.FloorNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Floor","block_type":"function","_x":3063,"_y":484,"in_parameters":[{"id":"7afe65d9-19ff-4b7a-a6f8-cf90b2c71c3e","name":"number","type":"System.Double","value":null,"assignment":"58cc0ac5-3ae1-4b56-b85f-270ea565fa07","assignment_node":"6ad3c998-1346-4643-ba70-5bda986661e3","value_is_reference":false}],"out_parameters":[{"id":"7089476c-5054-4f4c-a042-1a2df0716624","name":"value","type":"System.Double","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"92262ca7-fc16-493f-9bec-6bf0a89cbd7d","type":"*math.AddNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Add A + B","block_type":"function","_x":3386,"_y":492,"in_parameters":[{"id":"e00971da-c089-455c-8ff0-b90bf3db5c63","name":"a","type":"System.Double","value":null,"assignment":"7089476c-5054-4f4c-a042-1a2df0716624","assignment_node":"9d014de1-d97e-41ff-af62-defc2ca3bcf1","value_is_reference":false},{"id":"387f4da2-6122-49b9-8a95-3af3237184c4","name":"b","type":"System.Double","value":null,"assignment":"ba22f9a9-00e1-44a8-81b4-11a3cc72b2dd","assignment_node":"2735022b-0e7b-4a25-a5db-7e028792043b","value_is_reference":false}],"out_parameters":[{"id":"c9e41cfe-0e5a-4072-aac2-b54b2d4bdf62","name":"value","type":"System.Double","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"bf5bcf32-5aa8-4ebb-a104-513405eda1ce","type":"*math.CeilNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Ceiling","block_type":"function","_x":3654,"_y":494,"in_parameters":[{"id":"14af9a59-4d90-4481-91fd-8a5d25ebf3be","name":"number","type":"System.Double","value":null,"assignment":"c9e41cfe-0e5a-4072-aac2-b54b2d4bdf62","assignment_node":"92262ca7-fc16-493f-9bec-6bf0a89cbd7d","value_is_reference":false}],"out_parameters":[{"id":"12188199-30aa-43e8-b721-0058f07e5b0f","name":"value","type":"System.Double","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"632b1aff-32a2-4710-bc91-c8209ec0d645","type":"*math.AddNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Add A + B","block_type":"function","_x":4004,"_y":513,"in_parameters":[{"id":"48a0cf1a-3601-4861-b8ed-97bce5230ae3","name":"a","type":"System.Double","value":null,"assignment":"12188199-30aa-43e8-b721-0058f07e5b0f","assignment_node":"bf5bcf32-5aa8-4ebb-a104-513405eda1ce","value_is_reference":false},{"id":"31638b4f-a1dd-4375-bede-789d66089c4c","name":"b","type":"System.Double","value":null,"assignment":"fb8a55fd-ecb0-4eae-8436-cfa9823f66fd","assignment_node":"607ebbac-d2da-4a7a-93b8-3ff4024a79d1","value_is_reference":false}],"out_parameters":[{"id":"4b0fbf31-05c4-4a8d-a337-aac01ebab005","name":"value","type":"System.Double","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"607ebbac-d2da-4a7a-93b8-3ff4024a79d1","type":"*vars.DecimalNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Decimal","block_type":"variable","_x":3773,"_y":626,"in_parameters":[],"out_parameters":[{"id":"fb8a55fd-ecb0-4eae-8436-cfa9823f66fd","name":"value","type":"System.Double","value":"0.55235","assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"d181f4db-e20f-4a70-9070-7e3fe46a66cf","type":"*math.RoundNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Round","block_type":"function","_x":4309,"_y":520,"in_parameters":[{"id":"a0a73e8f-9360-431f-aa7d-d6d097c43042","name":"number","type":"System.Double","value":null,"assignment":"4b0fbf31-05c4-4a8d-a337-aac01ebab005","assignment_node":"632b1aff-32a2-4710-bc91-c8209ec0d645","value_is_reference":false},{"id":"775e2330-d67d-471a-96c4-f3d69fd3389b","name":"places","type":"System.Double","value":null,"assignment":"5ad07f78-f4ef-47f9-90c0-04848b5798d1","assignment_node":"454348fe-2bac-4efc-883b-6ee53c26274f","value_is_reference":false}],"out_parameters":[{"id":"ccaa8de0-5381-4a8b-9a90-a99216e6d7e5","name":"value","type":"System.Double","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"454348fe-2bac-4efc-883b-6ee53c26274f","type":"*vars.DecimalNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Decimal","block_type":"variable","_x":4050,"_y":629,"in_parameters":[],"out_parameters":[{"id":"5ad07f78-f4ef-47f9-90c0-04848b5798d1","name":"value","type":"System.Double","value":"4","assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"a7134284-bb63-44f2-ad98-41a5495965b3","type":"*math.TruncNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Truncate","block_type":"function","_x":4706,"_y":512,"in_parameters":[{"id":"65bae63a-5821-4b36-8a8d-af6dc2231a5d","name":"number","type":"System.Double","value":null,"assignment":"ccaa8de0-5381-4a8b-9a90-a99216e6d7e5","assignment_node":"d181f4db-e20f-4a70-9070-7e3fe46a66cf","value_is_reference":false},{"id":"9cc541fa-263d-470d-a822-df56ef099077","name":"precision","type":"System.Double","value":null,"assignment":"d1e35c18-17b4-442b-b47e-4a1169841ce3","assignment_node":"60176710-6665-4e74-a0f7-24398b7a2847","value_is_reference":false}],"out_parameters":[{"id":"1a656547-d65e-4d01-beef-4f7d3fcc3a43","name":"value","type":"System.Double","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"60176710-6665-4e74-a0f7-24398b7a2847","type":"*vars.DecimalNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Decimal","block_type":"variable","_x":4449,"_y":633,"in_parameters":[],"out_parameters":[{"id":"d1e35c18-17b4-442b-b47e-4a1169841ce3","name":"value","type":"System.Double","value":"2","assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"ab7e7f2f-e7c8-4fd0-a73d-0859d6c82f5a","type":"*math.PercentageDiffNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Percentage Difference","block_type":"function","_x":5170,"_y":511,"in_parameters":[{"id":"f6db1931-27b7-4bfc-9785-45878e6e55bf","name":"a","type":"System.Double","value":null,"assignment":"d857eccc-97cf-4958-af9b-a36bf49fc85f","assignment_node":"68ea2019-fb35-456a-b253-465e55c5bc7b","value_is_reference":false},{"id":"51345fba-9c21-42c0-a3c9-10ea7a0d9776","name":"b","type":"System.Double","value":null,"assignment":"1a656547-d65e-4d01-beef-4f7d3fcc3a43","assignment_node":"a7134284-bb63-44f2-ad98-41a5495965b3","value_is_reference":false}],"out_parameters":[{"id":"50625d3b-2458-4a8f-8a1b-82469840cfde","name":"value","type":"System.Double","value":null,"assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"68ea2019-fb35-456a-b253-465e55c5bc7b","type":"*vars.DecimalNode","out_node":null,"can_be_executed":false,"can_execute":false,"friendly_name":"Decimal","block_type":"variable","_x":4907,"_y":450,"in_parameters":[],"out_parameters":[{"id":"d857eccc-97cf-4958-af9b-a36bf49fc85f","name":"value","type":"System.Double","value":"100","assignment":"","assignment_node":"","value_is_reference":false}]},{"id":"339edee5-144c-4069-a4f4-2202350e3a67","type":"*console.PrintNode","out_node":null,"can_be_executed":true,"can_execute":true,"friendly_name":"Print","block_type":"function","_x":5529,"_y":331,"in_parameters":[{"id":"4d82a36f-33cb-4ccf-8c1d-489ce93cca74","name":"message","type":"System.String","value":null,"assignment":"50625d3b-2458-4a8f-8a1b-82469840cfde","assignment_node":"ab7e7f2f-e7c8-4fd0-a73d-0859d6c82f5a","value_is_reference":false}],"out_parameters":[]}],"comments":[]}`

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
	e.Run(context.Background())

	assert.Equal(t, result, "-40.45")
}
