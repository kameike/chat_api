#!/bin/sh

set -eu

swagger generate server -t swggen --exclude-main -f swagger.yml -A chat_api   --model-package="apimodel"

