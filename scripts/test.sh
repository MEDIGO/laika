#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

echo "mode: set" > combined_coverage.out

for d in $(go list ./... | grep -v vendor); do
    go test -cover -coverprofile=coverage.out $d
    if [ -f coverage.out ]; then
        cat coverage.out | grep -h -v "^mode:" >> combined_coverage.out
        rm coverage.out
    fi
done
