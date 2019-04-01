#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

if ! which glide > /dev/null 2>&1 ; then
  go get github.com/Masterminds/glide
fi

glide install
(cd dashboard && sudo apt-get remove cmdtest && sudo apt-get install yarn && yarn install)
