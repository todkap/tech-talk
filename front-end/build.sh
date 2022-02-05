#!/bin/sh
set -e

export DOCKER_ID_USER="todkap"
docker login

docker build --no-cache=true -t todkap/front-end:1.0.0 .

docker push todkap/front-end:1.0.0