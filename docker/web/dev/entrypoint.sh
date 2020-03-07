#!/bin/bash

set -e

usermod -u $WEB_UID www-data

#until nc -z sql 3306
#do
#    echo "Waiting for sql..."
#    sleep 1
#done

#su www-data bash -c "cd /var/www/script/migration && composer install --no-dev && ./vendor/bin/phinx migrate"
su www-data bash -c "composer install -d /var/www"

apachectl -D FOREGROUND
