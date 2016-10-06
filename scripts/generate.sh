#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

go-bindata -pkg schema -o store/schema/schema.go -ignore \.go store/schema/...
