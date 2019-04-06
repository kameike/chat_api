#!/bin/sh

set -eu

swagger generate server -t swggen --exclude-main -f swagger.yml -A chat_api   --model-package="apimodel"
swagger generate client -t swggen -f swagger.yml -A chat_api   --model-package="apimodel"

