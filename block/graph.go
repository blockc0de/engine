package block

import (
	"github.com/blockc0de/engine/attributes"
	"github.com/google/uuid"
)

type Graph struct {
	Hash            string
	Name            string
	Nodes           map[string]Node
	NodeList        []Node
	MemoryVariables map[string]interface{}
}

func NewGraph(hash, name string) *Graph {
	graph := Graph{
		Hash:            hash,
		Name:            name,
		Nodes:           make(map[string]Node),
		NodeList:        make([]Node, 0),
		MemoryVariables: make(map[string]interface{}),
	}
	return &graph
}

func (g *Graph) AddNode(node Node) {
	if node.Data().Id == "" {
		node.Data().Id = uuid.New().String()
	}
	g.Nodes[node.Data().Id] = node
	g.NodeList = append(g.NodeList, node)
}

func (g *Graph) FindNode(filter func(n Node) bool) Node {
	for _, node := range g.NodeList {
		if filter(node) {
			return node
		}
	}
	return nil
}

func (g *Graph) FindOutParameter(id string) *NodeParameter {
	for _, node := range g.NodeList {
		for _, parameter := range node.Data().OutParameters {
			if parameter.Id == id {
				return parameter
			}
		}
	}
	return nil
}

func (g *Graph) FindOutParameterNode(id string) Node {
	for _, node := range g.NodeList {
		for _, parameter := range node.Data().OutParameters {
			if parameter.Id == id {
				return node
			}
		}
	}
	return nil
}

func (g *Graph) GetEventNodes() []EventNode {
	nodes := make([]EventNode, 0)
	for _, node := range g.NodeList {
		if node.Data().IsEventNode {
			eventNode, ok := node.(EventNode)
			if ok {
				nodes = append(nodes, eventNode)
			}
		}
	}
	return nodes
}

func (g *Graph) GetFirstEntryPointNode() Node {
	for _, node := range g.NodeList {
		if node.Data().NodeBlockType == attributes.NodeTypeEnumEntryPoint {
			return node
		}
	}
	return nil
}
