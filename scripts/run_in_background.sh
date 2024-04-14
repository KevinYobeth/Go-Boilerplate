#!/bin/bash
set -e

export SSH_PRIVATE_KEY=$(cat ~/.ssh/id_rsa)

docker compose build
docker compose up -d

sleep 5

docker rm $(docker ps -a -f status=exited -f status=created -q)
docker image prune -f