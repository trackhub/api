#!/bin/bash

set -e

usermod -u $WEB_UID app

cp /usr/local/go/* /go-root-host/ -r
touch /go-root-host/go.mod
chown app /go-root-host/ -R
chown app /go -R

tail -f /var/log/*
