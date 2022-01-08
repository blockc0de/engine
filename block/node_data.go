package block

import (
	"reflect"

	"github.com/blockc0de/engine/attributes"
)

type NodeData struct {
	Id                       string
	Graph                    *Graph
	NodeType                 string
	FriendlyName             string
	NodeGroupName            string
	NodeBlockType            attributes.NodeTypeEnum
	NodeDescription          string
	BlockLimitPerGraph       int `json:",omitempty"`
	InParameters             NodeParameters
	OutParameters            NodeParameters
	OutNode                  Node
	IsEventNode              bool
	LastExecutionFrom        Node
	CanBeSerialized          bool
	NodeCycleLimit           int64
	LastCycleAt              int64
	CurrentTraceItem         *NodeTrace
	CustomTimeout            int64
	SpecialActionsAttributes []attributes.NodeSpecialActionAttribute
	IDEParameters            *attributes.NodeIDEParametersAttribute
}

func NewNodeData(id string, node Node, graph *Graph, nodeType string) *NodeData {
	n := new(NodeData)
	n.Id = id
	n.Graph = graph
	n.NodeType = nodeType
	n.CanBeSerialized = true
	n.SpecialActionsAttributes = make([]attributes.NodeSpecialActionAttribute, 0)

	customAttributes := node.GetCustomAttributes(reflect.TypeOf(attributes.NodeDefinition{}))
	if len(customAttributes) > 0 {
		nodeDefinition := customAttributes[0].(attributes.NodeDefinition)
		n.FriendlyName = nodeDefinition.FriendlyName
		n.NodeBlockType = nodeDefinition.NodeType
		n.NodeGroupName = nodeDefinition.GroupName
		n.BlockLimitPerGraph = nodeDefinition.BlockLimitPerGraph
	}

	customAttributes = node.GetCustomAttributes(reflect.TypeOf(attributes.NodeCycleLimit{}))
	if len(customAttributes) > 0 {
		nodeCycleLimit := customAttributes[0].(attributes.NodeCycleLimit)
		n.NodeCycleLimit = nodeCycleLimit.LimitPerCycle
	}

	customAttributes = node.GetCustomAttributes(reflect.TypeOf(attributes.NodeGraphDescription{}))
	if len(customAttributes) > 0 {
		nodeDescription := customAttributes[0].(attributes.NodeGraphDescription)
		n.NodeDescription = nodeDescription.Description
	}

	customAttributes = node.GetCustomAttributes(reflect.TypeOf(attributes.NodeTimeout{}))
	if len(customAttributes) > 0 {
		nodeTimeout := customAttributes[0].(attributes.NodeTimeout)
		n.CustomTimeout = nodeTimeout.CustomTimeout
	}

	customAttributes = node.GetCustomAttributes(reflect.TypeOf(attributes.NodeIDEParametersAttribute{}))
	if len(customAttributes) > 0 {
		nodeIdeParameters := customAttributes[0].(attributes.NodeIDEParametersAttribute)
		n.IDEParameters = &nodeIdeParameters
	}

	customAttributes = node.GetCustomAttributes(reflect.TypeOf(attributes.NodeSpecialActionAttribute{}))
	if len(customAttributes) > 0 {
		for _, attribute := range customAttributes {
			n.SpecialActionsAttributes = append(n.SpecialActionsAttributes, attribute.(attributes.NodeSpecialActionAttribute))
		}
	}

	n.InParameters = make(NodeParameters, 0)
	n.OutParameters = make(NodeParameters, 0)
	return n
}
