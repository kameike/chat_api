#!/bin/sh

set -eu

swagger generate server -t swggen --exclude-main -f swagger.yaml -A chat_api   --model-package="apimodel"
swagger generate client -t swggen -f swagger.yaml -A chat_api   --model-package="apimodel"

