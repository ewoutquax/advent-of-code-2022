#! /usr/bin/env bash

mkdir -p bin/

go build -o ./bin/solvePart1 ./cmd/part-1/main.go
go build -o ./bin/solvePart2 ./cmd/part-2/main.go

