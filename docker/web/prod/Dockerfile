FROM trackhub-web

CMD usermod -u $WEB_UID www-data && \
    su www-data bash -c "export APP_ENV=prod && composer install -d /var/www --no-dev -o" && \
    apachectl -D FOREGROUND
