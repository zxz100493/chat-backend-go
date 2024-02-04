#!/bin/bash
set -u
set -e
set -x
set -v

docker build -t apps:v1 -f ../Dockerfile ../ >/tmp/build.log

#delete dangling images
docker rmi -f $(docker images -q -f dangling=true)
