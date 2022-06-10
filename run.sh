#!/bin/sh

set -x

PRETTY=""
if [ -n "$1" ]
then
    PRETTY="-e JSON_PRETTY=1"
fi

docker run \
    -it \
    --rm \
    -p 8080:8080 \
    $PRETTY \
    steveww/webhook
