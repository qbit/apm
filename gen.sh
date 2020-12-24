#!/bin/sh

OS=$(uname -s | tr [:upper:] [:lower:])
go tool cgo -godefs gen_defs.go | gofmt > "apm_${OS}.go"
