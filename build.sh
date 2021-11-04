#!/bin/sh

set -e

IMAGE="steveww/webhook:latest"

docker build -t "$IMAGE" .

if [ "$1" = "push" ]
then
    docker push "$IMAGE"
fi
