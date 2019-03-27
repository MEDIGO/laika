#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

docker build --pull --rm medigo/laika .
