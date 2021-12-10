package main

import (
	"github.com/blockc0de/engine/loader"
	"io/ioutil"
)

func main() {
	schema, err := loader.ExportNodeSchema()
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("node_schema.json", schema, 0666)
	if err != nil {
		panic(err)
	}
}
