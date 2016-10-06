#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

commit=$(git rev-parse HEAD)

docker pull medigo/laika:$(commit)
aws ecs register-task-definition --family $(ECS_FAMILY) --container-definitions '$(shell ./ecs-container-definitions.sh)'
aws ecs update-service --service $(ECS_FAMILY) --task-definition $(ECS_FAMILY)
aws ecs wait services-stable --services $(ECS_FAMILY)
