#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

docker build --rm -t medigo/laika .
