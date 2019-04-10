#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

if ! which glide > /dev/null 2>&1 ; then
  go get github.com/Masterminds/glide
fi

glide install
(cd dashboard && sudo apt install curl && curl -sL https://deb.nodesource.com/setup_6.x | sudo bash - && sudo apt-get install -y nodejs && sudo apt-get install -y npm && sudo npm install)
