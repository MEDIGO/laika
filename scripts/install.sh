#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

bower install --allow-root
glide install
