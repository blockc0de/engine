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

	data, err := m.Inputs.PackValues(values)
	if err != nil {
		return nil, err
	}

	return append(m.ID, data...), nil
}

func (e *AbiEncoder) EncodeInputs(method string, jsonArray string) ([]byte, error) {
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

func (e *AbiEncoder) EncodeOutputs(method string, jsonArray string) ([]byte, error) {
	m, ok := e.abi.Methods[method]
	if !ok {
		return nil, errors.New("method not found")
	}

	values, err := e.parseParams(m, jsonArray)
	if err != nil {
		return nil, err
	}

	return m.Outputs.PackValues(values)
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
	case abi.TupleTy:
		return e.convertToTuple(value, t)
	case abi.ArrayTy:
		return e.convertToArray(value, t)
	case abi.SliceTy:
		return e.convertToSlice(value, t)
	case abi.BytesTy:
		return e.convertToBytes(value, t)
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

	var n int64
	var err error
	switch value := value.(type) {
	case string:
		n, err = strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
	case json.Number:
		n, err = strconv.ParseInt(string(value), 10, 64)
		if err != nil {
			return nil, err
		}
	default:
		return nil, ErrInvalidParamType(t.String())
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

	var ok bool
	var bn *big.Int
	switch value := value.(type) {
	case string:
		bn, ok = big.NewInt(0).SetString(value, 10)
		if !ok {
			return nil, ErrInvalidParamType(t.String())
		}
	case json.Number:
		bn, ok = big.NewInt(0).SetString(string(value), 10)
		if !ok {
			return nil, ErrInvalidParamType(t.String())
		}
	default:
		return nil, ErrInvalidParamType(t.String())
	}

	n := bn.Uint64()
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
		return bn, nil
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

func (e *AbiEncoder) convertToBytes(value interface{}, t abi.Type) (interface{}, error) {
	if t.T != abi.BytesTy {
		return nil, ErrInvalidParamType(t.String())
	}

	s, ok := value.(string)
	if !ok {
		return nil, ErrInvalidParamType(t.String())
	}
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
		s = s[2:]
	}

	return hex.DecodeString(s)
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

func (e *AbiEncoder) convertToTuple(value interface{}, t abi.Type) (interface{}, error) {
	if t.T != abi.TupleTy {
		return nil, ErrInvalidParamType(t.String())
	}

	slice, ok := value.([]interface{})
	if !ok {
		return nil, ErrInvalidParamType(t.String())
	}

	if len(slice) != len(t.TupleElems) {
		return nil, ErrInvalidParamType(t.String())
	}

	tuple := reflect.New(t.GetType())
	for idx := 0; idx < len(t.TupleElems); idx++ {
		elem, err := e.convertValue(slice[idx], *t.TupleElems[idx])
		if err != nil {
			return nil, err
		}

		field := tuple.Elem().Field(idx)
		field.Set(reflect.ValueOf(elem))
	}
	return tuple.Elem().Interface(), nil
}

func (e *AbiEncoder) convertToArray(value interface{}, t abi.Type) (interface{}, error) {
	if t.T != abi.ArrayTy {
		return nil, ErrInvalidParamType(t.String())
	}

	slice, ok := value.([]interface{})
	if !ok {
		return nil, ErrInvalidParamType(t.String())
	}

	goSlice := reflect.New(t.GetType())
	for idx := 0; idx < len(slice); idx++ {
		elem, err := e.convertValue(slice[idx], *t.Elem)
		if err != nil {
			return nil, err
		}
		goSlice.Elem().Index(idx).Set(reflect.ValueOf(elem))
	}
	return goSlice.Elem().Interface(), nil
}

func (e *AbiEncoder) convertToSlice(value interface{}, t abi.Type) (interface{}, error) {
	if t.T != abi.SliceTy {
		return nil, ErrInvalidParamType(t.String())
	}

	slice, ok := value.([]interface{})
	if !ok {
		return nil, ErrInvalidParamType(t.String())
	}

	goSlice := reflect.New(t.GetType())
	for idx := 0; idx < len(slice); idx++ {
		elem, err := e.convertValue(slice[idx], *t.Elem)
		if err != nil {
			return nil, err
		}
		goSlice.Elem().Set(reflect.Append(goSlice.Elem(), reflect.ValueOf(elem)))
	}
	return goSlice.Elem().Interface(), nil
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
	return goArray.Elem().Interface(), nil
}
