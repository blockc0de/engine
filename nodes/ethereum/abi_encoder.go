package ethereum

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	jsoniter "github.com/json-iterator/go"
)

var (
	jsonx = jsoniter.Config{
		EscapeHTML:             true,
		ValidateJsonRawMessage: true,
		UseNumber:              true,
	}.Froze()
)

func ErrInvalidParamType(s string) error {
	return fmt.Errorf("invalid type, want %s", s)
}

type AbiEncoder struct {
	abi abi.ABI
}

func NewAbiEncoder(a abi.ABI) AbiEncoder {
	return AbiEncoder{abi: a}
}

func (e *AbiEncoder) Encode(method string, jsonArray string) ([]byte, error) {
	m, ok := e.abi.Methods[method]
	if !ok {
		return nil, errors.New("method not found")
	}

	values, err := e.parseParams(m, jsonArray)
	if err != nil {
		return nil, err
	}

	return m.Inputs.PackValues(values)
}

func (e *AbiEncoder) parseParams(method abi.Method, jsonArray string) ([]interface{}, error) {
	var params []interface{}
	err := jsonx.Unmarshal([]byte(jsonArray), &params)
	if err != nil {
		return nil, err
	}

	if len(params) != len(method.Inputs) {
		return nil, fmt.Errorf("wrong number of params, want %d, got %d", len(method.Inputs), len(params))
	}

	values := make([]interface{}, 0, len(params))
	for idx, param := range params {
		value, err := e.convertValue(param, method.Inputs[idx].Type)
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}

	return values, nil
}

func (e *AbiEncoder) convertValue(value interface{}, t abi.Type) (interface{}, error) {
	switch t.T {
	case abi.IntTy:
		return e.convertToInt(value, t)
	case abi.UintTy:
		return e.convertToUInt(value, t)
	case abi.BoolTy:
		return e.convertToBool(value, t)
	case abi.StringTy:
		return e.convertToString(value, t)
	case abi.AddressTy:
		return e.convertToAddress(value, t)
	case abi.FixedBytesTy:
		return e.convertToFixedBytes(value, t)
	default:
		return nil, errors.New("type not supported")
	}
}

func (e *AbiEncoder) convertToInt(value interface{}, t abi.Type) (interface{}, error) {
	if t.T != abi.IntTy {
		return nil, ErrInvalidParamType(t.String())
	}

	number, ok := value.(json.Number)
	if !ok {
		return nil, ErrInvalidParamType(t.String())
	}

	n, err := strconv.ParseInt(string(number), 10, 64)
	if err != nil {
		return nil, err
	}

	switch t.Size {
	case 8:
		return int8(n), nil
	case 16:
		return int16(n), nil
	case 32:
		return int32(n), nil
	case 64:
		return n, nil
	default:
		return big.NewInt(n), nil
	}
}

func (e *AbiEncoder) convertToUInt(value interface{}, t abi.Type) (interface{}, error) {
	if t.T != abi.UintTy {
		return nil, ErrInvalidParamType(t.String())
	}

	number, ok := value.(json.Number)
	if !ok {
		return nil, ErrInvalidParamType(t.String())
	}

	b, ok := big.NewInt(0).SetString(string(number), 10)
	if !ok {
		return nil, ErrInvalidParamType(t.String())
	}

	n := b.Uint64()
	switch t.Size {
	case 8:
		return uint8(n), nil
	case 16:
		return uint16(n), nil
	case 32:
		return uint32(n), nil
	case 64:
		return n, nil
	default:
		return b, nil
	}
}

func (e *AbiEncoder) convertToBool(value interface{}, t abi.Type) (interface{}, error) {
	if t.T != abi.BoolTy {
		return nil, ErrInvalidParamType(t.String())
	}

	b, ok := value.(bool)
	if !ok {
		return nil, ErrInvalidParamType(t.String())
	}

	return b, nil
}

func (e *AbiEncoder) convertToString(value interface{}, t abi.Type) (interface{}, error) {
	if t.T != abi.StringTy {
		return nil, ErrInvalidParamType(t.String())
	}

	s, ok := value.(string)
	if !ok {
		return nil, ErrInvalidParamType(t.String())
	}

	return s, nil
}

func (e *AbiEncoder) convertToAddress(value interface{}, t abi.Type) (interface{}, error) {
	if t.T != abi.AddressTy {
		return nil, ErrInvalidParamType(t.String())
	}

	s, ok := value.(string)
	if !ok {
		return nil, ErrInvalidParamType(t.String())
	}

	return common.HexToAddress(s), nil
}

func (e *AbiEncoder) convertToFixedBytes(value interface{}, t abi.Type) (interface{}, error) {
	if t.T != abi.FixedBytesTy && t.Elem != nil {
		return nil, ErrInvalidParamType(t.String())
	}

	s, ok := value.(string)
	if !ok {
		return nil, ErrInvalidParamType(t.String())
	}
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
		s = s[2:]
	}

	data, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	fixedBytes := make([]byte, t.Size)
	copy(fixedBytes, data)

	goArray := reflect.New(t.GetType())
	for idx := 0; idx < t.Size; idx++ {
		goArray.Elem().Index(idx).Set(reflect.ValueOf(fixedBytes[idx]))
	}
	return goArray.Interface(), nil
}
