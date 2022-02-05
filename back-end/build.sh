#!/bin/sh
set -e

export DOCKER_ID_USER="todkap"
docker login

docker build --no-cache=true -t todkap/back-end-in-memory:1.0.0 .

docker push todkap/back-end-in-memory:1.0.0