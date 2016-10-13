#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

if ! which go-bindata > /dev/null 2>&1 ; then
  go get github.com/jteeuwen/go-bindata/...
fi

go-bindata -pkg schema -o store/schema/schema.go -ignore \.go store/schema/...
