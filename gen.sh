#!/bin/sh

go tool cgo -godefs gen_defs.go | gofmt > "apm_defs.go"
