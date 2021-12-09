package attributes

type NodeDefinition struct {
	NodeName           string
	FriendlyName       string
	NodeType           NodeTypeEnum
	GroupName          string
	BlockLimitPerGraph int
}
