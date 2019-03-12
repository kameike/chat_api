#!/bin/sh
set -eu

docker ps -q | xargs docker kill

docker run -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=test -d mysql:5.7
docker run -p 6379:6379 -d redis redis-server 

export CHAT_REDIS_ADDR=localhost:6379
export CHAT_REDIS_PASS=ihoge
export "CHAT_RDS_ADDR=root:password@tcp(localhost:3306)/test?parseTime=true"

sleep 10

go run cmd/migration/main.go -dir cmd/migration/ up

go test ./...
