#!/usr/bin/env bash

docker run --name traefik-proxy --rm -d \
    -p 8081:80 \
    -p 8080:8080 \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v $PWD/traefik.toml:/etc/traefik/traefik.toml \
    traefik:v1.7.4-alpine