#!/bin/sh
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./quiz_schema/quiz_schema.proto
