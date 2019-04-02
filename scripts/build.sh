#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/laika .
(cd dashboard && npm run build)
