package attributes

type NodeTypeEnum string

var (
	NodeTypeEnumFunction   NodeTypeEnum = "function"
	NodeTypeEnumEntryPoint NodeTypeEnum = "entry-point"
	NodeTypeEnumEvent      NodeTypeEnum = "event"
	NodeTypeEnumVariable   NodeTypeEnum = "variable"
	NodeTypeEnumCondition  NodeTypeEnum = "condition"
	NodeTypeEnumConnector  NodeTypeEnum = "connector"
	NodeTypeEnumDeployer   NodeTypeEnum = "deployer"
)
