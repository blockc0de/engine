package ethereum

import (
	"encoding/hex"
	"encoding/json"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/stretchr/testify/assert"
)

var TestABI abi.ABI
var TestAbiJSON = `[{"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"setAddress","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bool","name":"","type":"bool"}],"name":"setBool","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes","name":"","type":"bytes"}],"name":"setBytes","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"setBytes32","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"int16","name":"","type":"int16"}],"name":"setInt16","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"int256","name":"","type":"int256"}],"name":"setInt256","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"int32","name":"","type":"int32"}],"name":"setInt32","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"int64","name":"","type":"int64"}],"name":"setInt64","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"int8","name":"","type":"int8"}],"name":"setInt8","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"","type":"string"}],"name":"setString","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string[2]","name":"","type":"string[2]"}],"name":"setStringArray","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string[]","name":"","type":"string[]"}],"name":"setStringSlice","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"components":[{"internalType":"uint256","name":"amount","type":"uint256"},{"internalType":"address[]","name":"addresses","type":"address[]"}],"internalType":"struct Test.Funder","name":"","type":"tuple"}],"name":"setStruct","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint16","name":"","type":"uint16"}],"name":"setUint16","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"setUint256","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint32","name":"","type":"uint32"}],"name":"setUint32","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint64","name":"","type":"uint64"}],"name":"setUint64","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint8","name":"","type":"uint8"}],"name":"setUint8","outputs":[],"stateMutability":"nonpayable","type":"function"}]`

func init() {
	var err error
	TestABI, err = abi.JSON(strings.NewReader(TestAbiJSON))
	if err != nil {
		panic(err)
	}
}

func unpackToJSON(t *testing.T, method string, data []byte) string {
	v, err := TestABI.Methods[method].Inputs.Unpack(data)
	assert.Nil(t, err)

	data, err = json.Marshal(v)
	assert.Nil(t, err)

	return string(data)
}

func TestAbiEncoder(t *testing.T) {
	encoder := NewAbiEncoder(TestABI)

	data, err := encoder.EncodeInputs("setInt8", `["88"]`)
	assert.Nil(t, err)
	assert.Equal(t, unpackToJSON(t, "setInt8", data), `[88]`)

	data, err = encoder.EncodeInputs("setInt16", `[1024]`)
	assert.Nil(t, err)
	assert.Equal(t, unpackToJSON(t, "setInt16", data), `[1024]`)

	data, err = encoder.EncodeInputs("setInt32", `[1024000]`)
	assert.Nil(t, err)
	assert.Equal(t, unpackToJSON(t, "setInt32", data), `[1024000]`)

	data, err = encoder.EncodeInputs("setInt64", `[1024000]`)
	assert.Nil(t, err)
	assert.Equal(t, unpackToJSON(t, "setInt64", data), `[1024000]`)

	data, err = encoder.EncodeInputs("setInt256", `[1024000]`)
	assert.Nil(t, err)
	assert.Equal(t, unpackToJSON(t, "setInt256", data), `[1024000]`)

	data, err = encoder.EncodeInputs("setUint8", `["88"]`)
	assert.Nil(t, err)
	assert.Equal(t, unpackToJSON(t, "setUint8", data), `[88]`)

	data, err = encoder.EncodeInputs("setUint16", `[1024]`)
	assert.Nil(t, err)
	assert.Equal(t, unpackToJSON(t, "setUint16", data), `[1024]`)

	data, err = encoder.EncodeInputs("setUint32", `[1024000]`)
	assert.Nil(t, err)
	assert.Equal(t, unpackToJSON(t, "setUint32", data), `[1024000]`)

	data, err = encoder.EncodeInputs("setUint64", `[1024000]`)
	assert.Nil(t, err)
	assert.Equal(t, unpackToJSON(t, "setUint64", data), `[1024000]`)

	data, err = encoder.EncodeInputs("setUint256", `[1024000]`)
	assert.Nil(t, err)
	assert.Equal(t, unpackToJSON(t, "setUint256", data), `[1024000]`)

	data, err = encoder.EncodeInputs("setBool", `[true]`)
	assert.Nil(t, err)
	assert.Equal(t, unpackToJSON(t, "setBool", data), `[true]`)

	data, err = encoder.EncodeInputs("setBytes32", `["0x42797465733332"]`)
	assert.Nil(t, err)
	assert.Equal(t, hex.EncodeToString(data), "4279746573333200000000000000000000000000000000000000000000000000")

	data, err = encoder.EncodeInputs("setBytes", `["0x42797465733332"]`)
	assert.Nil(t, err)
	assert.Equal(t, hex.EncodeToString(data), "000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000074279746573333200000000000000000000000000000000000000000000000000")

	data, err = encoder.EncodeInputs("setString", `["hello"]`)
	assert.Nil(t, err)
	assert.Equal(t, unpackToJSON(t, "setString", data), `["hello"]`)

	data, err = encoder.EncodeInputs("setStringSlice", `[["hello","world"]]`)
	assert.Nil(t, err)
	assert.Equal(t, unpackToJSON(t, "setStringSlice", data), `[["hello","world"]]`)

	data, err = encoder.EncodeInputs("setStringArray", `[["hello","world"]]`)
	assert.Nil(t, err)
	assert.Equal(t, unpackToJSON(t, "setStringArray", data), `[["hello","world"]]`)

	data, err = encoder.EncodeInputs("setAddress", `["0xe66666657a9b513288bf5b58be33a405f04f1749"]`)
	assert.Nil(t, err)
	assert.Equal(t, unpackToJSON(t, "setAddress", data), `["0xe66666657a9b513288bf5b58be33a405f04f1749"]`)

	data, err = encoder.EncodeInputs("setStruct", `[[1024, ["0xe66666657a9b513288bf5b58be33a405f04f1749"]]]`)
	assert.Nil(t, err)
	assert.Equal(t, unpackToJSON(t, "setStruct", data), `[{"amount":1024,"addresses":["0xe66666657a9b513288bf5b58be33a405f04f1749"]}]`)
}
