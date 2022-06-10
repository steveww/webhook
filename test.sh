#!/bin/sh

curl \
    -X POST \
    -d '{"foo": "bar"}' \
    -H 'Bar: Foo' \
    -H 'Content-Type: application/json' \
    http://localhost:8080/webhook
