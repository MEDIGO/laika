#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

if [[ -z ${CIRCLE_SHA1:-} ]]; then
  COMMIT=$(git rev-parse HEAD)
else
  COMMIT=${CIRCLE_SHA1}
fi

docker tag quay.io/medigo/laika:latest quay.io/medigo/laika:${COMMIT}
docker login -e ${DOCKER_EMAIL} -u ${DOCKER_USER} -p ${DOCKER_PASS}
docker push quay.io/medigo/laika:latest
docker push quay.io/medigo/laika:${COMMIT}
