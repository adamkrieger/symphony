#!/usr/bin/env bash
docker run --rm -it \
    --name symph-restapi-mono \
    -p 8001:8001 -e RUNPORT=8001 symph-restapi