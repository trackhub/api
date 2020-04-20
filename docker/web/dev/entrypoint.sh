#!/bin/bash

set -e

usermod -u $WEB_UID www-data

su www-data bash -c "composer install -d /var/www"

apachectl -D FOREGROUND
