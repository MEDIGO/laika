#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail
pwd
ls
docker build --pull --rm -t medigo/laika . -f Dockerfile
