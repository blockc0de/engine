#!/bin/bash

go test ./interop
go test ./compress
go test ./nodes/math
go test ./nodes/text
go test ./nodes/time
go test ./nodes/ethereum
go test ./nodes/ethereum/web3util
go test ./test
