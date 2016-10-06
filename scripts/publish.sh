#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

COMMIT=$(git rev-parse HEAD)

docker tag medigo/laika:latest medigo/laika:${COMMIT}
docker login -e ${DOCKER_EMAIL} -u ${DOCKER_USER} -p ${DOCKER_PASS}
docker push medigo/laika:latest
docker push medigo/laika:${COMMIT}
