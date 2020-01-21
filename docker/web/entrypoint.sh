#!/bin/bash

set -e

usermod -u ${WEB_UID} www-data

until nc -z sql 3306
do
    echo "Waiting for MariaDB..."
    sleep 1
done

apachectl -D FOREGROUND
