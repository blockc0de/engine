package block

import "context"

type NodeExecutor interface {
	AppendLog(string, string)
	GetUserData(key string) interface{}
	SetUserData(key string, data interface{})
	NextNode(ctx context.Context, node Node) bool
	ExecuteNode(ctx context.Context, node Node, executedFromNode Node) bool
	Stop()
}
