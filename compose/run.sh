#!/usr/bin/env sh
ko build -t devenv -LB ../.
hwclock -s
docker compose pull --ignore-pull-failures
docker compose up