#!/bin/bash

docker rm -f mysql

docker run \
    --name mysql \
    -d \
    -e MYSQL_DATABASE=lyrics \
    -e MYSQL_PASSWORD=lyrics \
    -e MYSQL_ROOT_PASSWORD=password \
    -e MYSQL_USER=lyrics \
    -p 3306:3306 \
    mysql:latest