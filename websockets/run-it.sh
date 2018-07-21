#!/usr/bin/env bash
docker run --rm -it \
    --name confcallmono \
    -p 8053:8053 -e RUNPORT=8053 confcall
