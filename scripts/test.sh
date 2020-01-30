#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

docker-compose run laika go test -v ./... -cover -coverprofile=combined_coverage.out
