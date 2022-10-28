#!/bin/bash

set -e

# usermod -u $WEB_UID app
cp /usr/local/go/* /go-root-host/ -r
chown $WEB_UID /go-root-host/ -R

tail -f /var/log/*
