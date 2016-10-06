#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

go run main.go migrate
