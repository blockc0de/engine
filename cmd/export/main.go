package main

import (
	"github.com/blockc0de/engine/interop"
	"io/ioutil"
)

func main() {
	schema, err := interop.ExportNodeSchema()
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("node_schema.json", schema, 0666)
	if err != nil {
		panic(err)
	}
}
