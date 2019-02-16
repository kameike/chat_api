#!/bin/sh

set -eu

swagger generate server -f swagger.yml -A chat_api   --model-package="apimodel"

