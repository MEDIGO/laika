#!/bin/sh -e

cat << EOM
[
    {
       "environment": [
           {
               "name": "LAIKA_MYSQL_USERNAME",
               "value": "$ECS_LAIKA_MYSQL_USERNAME"
           },
           {
               "name": "LAIKA_MYSQL_PASSWORD",
               "value": "$ECS_LAIKA_MYSQL_PASSWORD"
           },
           {
               "name": "LAIKA_MYSQL_HOST",
               "value": "$ECS_LAIKA_MYSQL_HOST"
           }
       ],
       "name": "laika",
       "mountPoints": [],
       "image": "quay.io/medigo/laika:$(git rev-parse HEAD)",
       "cpu": 128,
       "portMappings": [
           {
               "protocol": "tcp",
               "containerPort": 8000,
               "hostPort": $ECS_LAIKA_PORT
           }
       ],
       "command": [
           "laika"
       ],
       "memory": 128,
       "essential": true,
       "volumesFrom": []
   }
]
EOM
