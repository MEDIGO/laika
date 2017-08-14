#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

if ! which rerun > /dev/null 2>&1 ; then
  go get github.com/ivpusic/rerun
fi

npm run watch &
rerun -a run -i dashboard,node_modules,bin,public,vendor,.git
