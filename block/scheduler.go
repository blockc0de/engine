package block

import "context"

type NodeScheduler interface {
	AddCycle(node StartNode, parameters NodeParameters)
	NextNode(ctx context.Context, node Node) bool
	ExecuteNode(ctx context.Context, node Node, executedFromNode Node) bool
	AppendLog(string, string)
	Stop()
}
