#!/bin/bash
set -u
set -x
set -v

containerName=test-apps

existsContainer=$(docker ps -a --format "table {{.Names}}" | grep -v "NAMES" | grep ${containerName})

if [ $existsContainer = $containerName ]; then
    echo "${containerName} container exists"
    docker rm -f ${containerName}

else
    echo "${containerName} not running"

fi

set -e

docker run -d --name ${containerName} -p 8888:8888 -p 8080:8080 apps:v1
