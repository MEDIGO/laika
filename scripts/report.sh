#!/bin/sh

set -o errexit
set -o pipefail

if [[ -x "${COVERALLS_TOKEN}" ]]; then
  goveralls -coverprofile=combined_coverage.out -service=circle-ci -repotoken=$COVERALLS_TOKEN
fi
