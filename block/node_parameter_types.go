package block

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

type (
	NodeParameterBool    bool
	NodeParameterString  string
	NodeParameterDecimal struct{ decimal.Decimal }
)

type NodeParameterTypeEnum string

var (
	NodeParameterTypeEnumNode    NodeParameterTypeEnum = "NodeBlock.Engine.Node"
	NodeParameterTypeEnumBool    NodeParameterTypeEnum = "System.Boolean"
	NodeParameterTypeEnumString  NodeParameterTypeEnum = "System.String"
	NodeParameterTypeEnumDecimal NodeParameterTypeEnum = "System.Double"
	NodeParameterTypeEnumObject  NodeParameterTypeEnum = "System.Object"
	NodeParameterTypeEnumArray   NodeParameterTypeEnum = "System.Collections.Generic.List"
	NodeParameterTypeEnumUnknown NodeParameterTypeEnum = "Unknown"
)

func GetNodeParameterTypeEnum(v interface{}) NodeParameterTypeEnum {
	switch v.(type) {
	case Node:
		return NodeParameterTypeEnumNode
	case NodeParameterBool:
		return NodeParameterTypeEnumBool
	case NodeParameterString:
		return NodeParameterTypeEnumString
	case NodeParameterDecimal:
		return NodeParameterTypeEnumDecimal
	default:
		return NodeParameterTypeEnumUnknown
	}
}

func (b NodeParameterBool) MarshalJSON() ([]byte, error) {

	if b {
		return []byte(`"true"`), nil
	}
	return []byte(`"false"`), nil
}

func (b *NodeParameterBool) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, []byte(`"0"`)) == 0 {
		*b = false
		return nil
	}
	if bytes.Compare(data, []byte(`"1"`)) == 0 {
		*b = true
		return nil
	}
	if bytes.Compare(data, []byte(`"true"`)) == 0 {
		*b = true
		return nil
	}
	if bytes.Compare(data, []byte(`"false"`)) == 0 {
		*b = false
		return nil
	}
	return errors.New("invalid boolean value")
}

func (d NodeParameterDecimal) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, d.String())), nil
}

func (d *NodeParameterDecimal) UnmarshalJSON(data []byte) error {
	var v decimal.Decimal
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*d = NodeParameterDecimal{v}
	return nil
}

type NodeParameterConverter struct {
}

func (NodeParameterConverter) ToBool(val interface{}) (bool, bool) {
	switch val.(type) {
	case bool:
		return val.(bool), true
	case NodeParameterBool:
		return bool(val.(NodeParameterBool)), true
	case NodeParameterString:
		s := val.(NodeParameterString)
		if s == "0" || s == "false" {
			return false, true
		}
		return true, true
	case NodeParameterDecimal:
		d := val.(NodeParameterDecimal)
		return !d.IsZero(), true
	default:
		return false, false
	}
}

func (NodeParameterConverter) ToString(val interface{}) (string, bool) {
	switch val.(type) {
	case string:
		return val.(string), true
	case Node:
		node := val.(Node)
		if node != nil {
			return node.Data().Id, true
		}
		return "", false
	case NodeParameterBool:
		b := val.(NodeParameterBool)
		if b {
			return "true", true
		}
		return "false", true
	case NodeParameterString:
		return string(val.(NodeParameterString)), true
	case NodeParameterDecimal:
		d := val.(NodeParameterDecimal)
		return d.String(), true
	default:
		return "", false
	}
}

func (NodeParameterConverter) ToDecimal(val interface{}) (decimal.Decimal, bool) {
	switch val.(type) {
	case float64:
		return decimal.NewFromFloat(val.(float64)), true
	case NodeParameterBool:
		b := val.(NodeParameterBool)
		if b {
			return decimal.NewFromInt(1), true
		}
		return decimal.Zero, true
	case NodeParameterString:
		s := val.(NodeParameterString)
		d, err := decimal.NewFromString(string(s))
		if err != nil {
			return decimal.Zero, false
		}
		return d, true
	case NodeParameterDecimal:
		return val.(NodeParameterDecimal).Decimal, true
	default:
		return decimal.Zero, false
	}
}
