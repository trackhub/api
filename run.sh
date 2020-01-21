#!/usr/bin/env bash

set -e


[[ $UID == 0 ]] && export WEB_UID=1000 || export WEB_UID=$UID

docker-compose up --build
