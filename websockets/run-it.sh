#!/usr/bin/env bash
docker run --rm -it \
    --name confcallmono \
    -e RUNPORT=8053 \
    -p 8053:8053 -e RUNPORT=8053 confcall
