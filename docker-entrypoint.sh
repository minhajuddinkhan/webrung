#!/bin/sh

echo "running seed..."
go run bin/seed.go
echo "starting server..."
go run server/webrung/main.go