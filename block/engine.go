package block

import "context"

type Engine interface {
	AddCycle(node StartNode, parameters NodeParameters)
	CurrentCycle() *GraphExecutionCycle
	NextNode(ctx context.Context, node Node) bool
	ExecuteNode(ctx context.Context, node Node, executedFromNode Node) bool
	AppendLog(string, string)
	Stop()
}
