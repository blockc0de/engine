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
	NodeParameterTypeEnumObject  NodeParameterTypeEnum = "System.Object"
	NodeParameterTypeEnumDecimal NodeParameterTypeEnum = "Decimal"
	NodeParameterTypeEnumArray   NodeParameterTypeEnum = "List"
	NodeParameterTypeEnumMapping NodeParameterTypeEnum = "Mapping"
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
	if bytes.Equal(data, []byte(`"0"`)) {
		*b = false
		return nil
	}
	if bytes.Equal(data, []byte(`"1"`)) {
		*b = true
		return nil
	}
	if bytes.Equal(data, []byte(`"true"`)) {
		*b = true
		return nil
	}
	if bytes.Equal(data, []byte(`"false"`)) {
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
	switch val := val.(type) {
	case bool:
		return val, true
	case NodeParameterBool:
		return bool(val), true
	case NodeParameterString:
		if val == "0" || val == "false" {
			return false, true
		}
		return true, true
	case NodeParameterDecimal:
		return !val.IsZero(), true
	default:
		return false, false
	}
}

func (NodeParameterConverter) ToString(val interface{}) (string, bool) {
	switch val := val.(type) {
	case string:
		return val, true
	case Node:
		if val != nil {
			return val.Data().Id, true
		}
		return "", false
	case NodeParameterBool:
		if val {
			return "true", true
		}
		return "false", true
	case NodeParameterString:
		return string(val), true
	case NodeParameterDecimal:
		return val.String(), true
	default:
		return "", false
	}
}

func (NodeParameterConverter) ToDecimal(val interface{}) (decimal.Decimal, bool) {
	switch val := val.(type) {
	case float64:
		return decimal.NewFromFloat(val), true
	case NodeParameterBool:
		if val {
			return decimal.NewFromInt(1), true
		}
		return decimal.Zero, true
	case NodeParameterString:
		d, err := decimal.NewFromString(string(val))
		if err != nil {
			return decimal.Zero, false
		}
		return d, true
	case NodeParameterDecimal:
		return val.Decimal, true
	default:
		return decimal.Zero, false
	}
}
