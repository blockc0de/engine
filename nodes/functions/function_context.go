package functions

type FunctionContext struct {
	Node           *FunctionNode
	ReturnValues   map[string]interface{}
	CallParameters map[string]interface{}
}

func NewFunctionContext(node *FunctionNode) *FunctionContext {
	ctx := FunctionContext{
		Node:           node,
		ReturnValues:   make(map[string]interface{}),
		CallParameters: make(map[string]interface{}),
	}
	return &ctx
}
