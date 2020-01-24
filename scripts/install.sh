#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

go mod download
(cd dashboard && sudo apt install curl && curl -sL https://deb.nodesource.com/setup_6.x | sudo bash - && sudo apt-get install -y nodejs && sudo apt-get install -y npm && sudo npm install)

