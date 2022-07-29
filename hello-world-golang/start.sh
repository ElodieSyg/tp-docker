#!/bin/sh
./docker-gs-ping &
nginx -g 'daemon off;'