#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

if [[ -z ${CIRCLE_SHA1:-} ]]; then
  COMMIT=$(git rev-parse HEAD)
else
  COMMIT=${CIRCLE_SHA1}
fi

docker tag ${DOCKER_USER}/laika:latest ${DOCKER_USER}/laika:${COMMIT}
docker login -u ${DOCKER_USER} -p ${DOCKER_PASS}
docker push ${DOCKER_USER}/laika:latest
docker push ${DOCKER_USER}/laika:${COMMIT}
