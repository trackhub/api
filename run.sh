#!/bin/bash

set -e

# If the script is run with sudo, UID is 0. This is an issue when running
# "usermod -u $WEB_UID www-data" in the web container.
# In this case assign WEB_UID to 1000
[[ $UID == 0 ]] && export WEB_UID=1000 || export WEB_UID=$UID

if [[ "$1" == "dev" ]]; then
    rm docker/web/dev/goroot -rf || echo "Unable to delete goroot"
    echo "dev"
fi

docker build --tag=trackhub-api-web ./docker/web/
docker-compose -p track build

if [[ "$1" == "prod" ]]; then
  echo "Not implemented"
  exit 1
else
  if [[ "$1" == "dev" ]]; then
    docker-compose -p track -f docker-compose.yml -f docker-compose-dev.yml build
    docker-compose -p track -f docker-compose.yml -f docker-compose-dev.yml up
  fi
fi
