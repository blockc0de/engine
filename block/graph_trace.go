package block

import (
	"sort"
	"strconv"
	"strings"
)

type GraphTrace struct {
	error error
	stack []*NodeTrace
}

func NewGraphTrace() *GraphTrace {
	return &GraphTrace{stack: make([]*NodeTrace, 0)}
}

func (trace *GraphTrace) AppendTrace(node Node) *NodeTrace {
	item := NewNodeTrace(node)
	trace.stack = append(trace.stack, item)
	return item
}

func (trace *GraphTrace) SetException(err error) {
	trace.error = err
}

func (trace *GraphTrace) String() string {
	var errCount int
	var totalExecutionTime int64
	var executionSuccessText = "FALSE"
	traces := make([]string, 0, len(trace.stack))

	for _, item := range trace.stack {
		if item.ExecutionError != nil {
			errCount += 1
		}
		totalExecutionTime += item.ExecutionTime
		traces = append(traces, item.String())
	}

	if errCount == 0 {
		executionSuccessText = "TRUE"
	}

	sort.Reverse(sort.StringSlice(traces))
	return "Total execution time : " + strconv.FormatInt(totalExecutionTime, 10) + "ms\n" +
		"Execution success : " + executionSuccessText + "\n" +
		strings.Join(traces, "\n") + "\n" +
		"Stack count: " + strconv.Itoa(len(trace.stack))
}
