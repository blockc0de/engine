package ethereum

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	logDecoderNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "LogDecoderNode", FriendlyName: "Ethereum Log Decoder", NodeType: attributes.NodeTypeEnumCondition, GroupName: "Blockchain.Ethereum", BlockLimitPerGraph: -1}}
	logDecoderNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Decoding event logs from ethereum transactions"}}
)

type eventLog struct {
	Event  string        `json:"event,omitempty"`
	Inputs []interface{} `json:"inputs,omitempty"`
}

type LogDecoderNode struct {
	block.NodeBase
}

func NewLogDecoderNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(LogDecoderNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	eventLog, err := block.NewNodeParameter(node, "eventLog", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(eventLog)

	abi, err := block.NewNodeParameter(node, "abi", block.NodeParameterTypeEnumObject, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(abi)

	result, err := block.NewNodeParameter(node, "result", block.NodeParameterTypeEnumString, false, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(result)

	return node, err
}

func (n *LogDecoderNode) CanExecute() bool {
	return true
}

func (n *LogDecoderNode) CanBeExecuted() bool {
	return true
}

func (n *LogDecoderNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return logDecoderNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return logDecoderNodeGraphDescription
	default:
		return nil
	}
}

func (n *LogDecoderNode) OnExecution(ctx context.Context, scheduler block.NodeScheduler) error {
	value := n.Data().InParameters.Get("abi").ComputeValue()
	if value == nil {
		return block.ErrInvalidParameter{Name: "abi"}
	}
	abiInstance, ok := value.(*abi.ABI)
	if !ok {
		return block.ErrInvalidParameter{Name: "abi"}
	}

	var converter block.NodeParameterConverter
	s, ok := converter.ToString(n.Data().InParameters.Get("eventLog").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "eventLog"}
	}

	var log types.Log
	err := log.UnmarshalJSON([]byte(s))
	if err != nil {
		return block.ErrInvalidParameter{Name: "eventLog"}
	}

	result, err := n.decodeEventLog(abiInstance, &log)
	data, err := json.Marshal(result)
	if err != nil {
		return err
	}

	n.Data().OutParameters.Get("result").Value = block.NodeParameterString(data)

	return nil
}

func (n *LogDecoderNode) decodeEventLog(abiInstance *abi.ABI, log *types.Log) (eventLog, error) {
	var event *abi.Event
	for _, ev := range abiInstance.Events {
		if bytes.Equal(ev.ID[:], log.Topics[0][:]) {
			event = &ev
			break
		}
	}
	if event == nil {
		return eventLog{}, errors.New("mismatch")
	}

	inputs, err := abiInstance.Unpack(event.Name, log.Data)
	if err != nil {
		return eventLog{}, err
	}

	nextInput := 0
	nextIndex := 1
	result := make([]interface{}, len(event.Inputs))
	for idx, input := range event.Inputs {
		if !input.Indexed {
			result[idx] = inputs[nextInput]
			nextInput++
		} else {
			switch input.Type.String() {
			case "address":
				result[idx] = common.BytesToAddress(log.Topics[nextIndex].Bytes())
			default:
				result[idx] = log.Topics[nextIndex]
			}
			nextIndex++
		}
	}

	return eventLog{Event: event.RawName, Inputs: result}, nil
}
