package block

import (
	"strconv"
	"strings"
)

type NodeTrace struct {
	node           Node
	nodeId         string
	parameters     map[string]string
	ExecutionTime  int64
	ExecutionError error
}

func NewNodeTrace(node Node) *NodeTrace {
	trace := NodeTrace{
		node:           node,
		nodeId:         node.Data().Id,
		parameters:     make(map[string]string),
		ExecutionTime:  0,
		ExecutionError: nil,
	}

	for _, parameter := range node.Data().InParameters {
		if parameter == nil {
			trace.parameters[parameter.Name] = "NULL"
		} else {
			var converter NodeParameterConverter
			value, ok := converter.ToString(parameter.Value)
			if ok {
				trace.parameters[parameter.Name] = value
			}
		}
	}
	return &trace
}

func (trace *NodeTrace) String() string {
	lines := make([]string, 0)
	lines = append(lines, "at "+trace.node.Data().NodeType+" "+strconv.FormatInt(trace.ExecutionTime, 10)+"ms ("+trace.nodeId+")")
	if trace.ExecutionError != nil {
		lines = append(lines, "Exception occurred : "+trace.ExecutionError.Error())
	}
	for key, val := range trace.parameters {
		lines = append(lines, "-- "+key+" = "+val)
	}
	return strings.Join(lines, "\n")
}
