package interop

import (
	"github.com/blockc0de/engine/block"
	jsoniter "github.com/json-iterator/go"
)

// LoadGraph import graph from JSON string.
func LoadGraph(graphJson []byte) (*block.Graph, error) {
	var graphSchema GraphSchema
	err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(graphJson, &graphSchema)
	if err != nil {
		return nil, err
	}

	graph := block.NewGraph(graphSchema.ProjectId, graphSchema.Name)

	// Load nodes
	for _, nodeSchema := range graphSchema.Nodes {
		node, err := NewNode(nodeSchema.Type, nodeSchema.Id, graph)
		if err != nil {
			return nil, err
		}
		graph.AddNode(node)
	}

	// Load assignments and exec direction
	for _, nodeSchema := range graphSchema.Nodes {
		node, ok := graph.Nodes[nodeSchema.Id]
		if !ok {
			continue
		}

		if nodeSchema.OutNode != nil {
			if outNode, ok := graph.Nodes[*nodeSchema.OutNode]; ok {
				node.Data().OutNode = outNode
			}
		}

		for _, parameter := range nodeSchema.InParameters {
			nodeParam := node.Data().InParameters.Get(parameter.Name)
			if nodeParam != nil {
				nodeParam.Id = parameter.Id
				nodeParam.Value = parameter.Value
			}
		}

		for _, parameter := range nodeSchema.OutParameters {
			nodeParam := node.Data().OutParameters.Get(parameter.Name)
			if nodeParam == nil {
				continue
			}

			nodeParam.Id = parameter.Id
			if parameter.ValueIsReference && parameter.Value != nil {
				id, ok := parameter.Value.(string)
				if ok && len(id) > 0 {
					if reference, ok := graph.Nodes[id]; ok {
						nodeParam.Value = reference
					}
				} else {
					nodeParam.Value = parameter.Value
				}
			} else {
				nodeParam.Value = parameter.Value
			}
		}
	}

	for _, nodeSchema := range graphSchema.Nodes {
		node, ok := graph.Nodes[nodeSchema.Id]
		if !ok {
			continue
		}

		for _, parameter := range nodeSchema.InParameters {
			nodeParam := node.Data().InParameters.Get(parameter.Name)
			if nodeParam == nil {
				continue
			}

			if parameter.Assignment != "" {
				nodeParam.Assignments = graph.FindOutParameter(parameter.Assignment)
			}
		}

		for _, parameter := range nodeSchema.OutParameters {
			nodeParam := node.Data().OutParameters.Get(parameter.Name)
			if nodeParam == nil {
				continue
			}

			if parameter.Assignment != "" {
				nodeParam.Assignments = graph.FindOutParameter(parameter.Assignment)
			}
		}
	}

	return graph, nil
}

// ExportGraph export graph to JSON format.
func ExportGraph(graph *block.Graph) ([]byte, error) {
	schema := NewGraphSchema(graph)
	return schema.Export()
}

// ExportNodeSchema export schema of all nodes to JSON format.
func ExportNodeSchema() ([]byte, error) {
	nodes := make([]block.Node, 0, len(nodeCreators))
	for _, creator := range nodeCreators {
		node, err := creator.Creator("", nil)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	array := make([]map[string]interface{}, 0, len(nodes))
	for _, node := range nodes {
		data, err := json.Marshal(node)
		if err != nil {
			return nil, err
		}

		var mapper map[string]interface{}
		if err := json.Unmarshal(data, &mapper); err != nil {
			return nil, err
		}

		nodeData := mapper["NodeData"].(map[string]interface{})

		for key, val := range nodeData {
			mapper[key] = val
		}
		delete(mapper, "NodeData")

		mapper["CanExecute"] = node.CanExecute()
		mapper["CanBeExecuted"] = node.CanBeExecuted()

		array = append(array, mapper)
	}
	return json.Marshal(array)
}
