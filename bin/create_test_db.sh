#!/bin/sh
set -eu

docker run -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=test -d mysql:5.7
docker run -p 6379:6379 -d redis redis-server 
