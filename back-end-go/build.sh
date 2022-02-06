#!/bin/sh
set -e

export DOCKER_ID_USER="todkap"
docker login

docker build --no-cache=true -t todkap/back-end:1.0.1 .

docker push todkap/back-end:1.0.1