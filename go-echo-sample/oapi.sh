#!/bin/bash

oapi-codegen -generate "types" -package openapi sample-if/go-echo-api.yaml > ./openapi/types.gen.go
oapi-codegen -generate "server" -package openapi sample-if/go-echo-api.yaml > ./openapi/server.gen.go

