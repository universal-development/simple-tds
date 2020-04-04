#!/usr/bin/env bash

set -x

export REGISTRY=${REGISTRY:-denis256/simple-tds}
export TAG=${TAG:-0.0.1}

chmod +x ./simple-tds

docker build . -t "${REGISTRY}:${TAG}"
docker push "${REGISTRY}:${TAG}"
