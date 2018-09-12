#!/bin/bash

script_dir=$(dirname $0)
cd $script_dir

set -o verbose

TAG=$1

if [[ -z $TAG ]]
then
	TAG="latest"
fi

docker build -t liftedkilt/foothillnorthchur.ch-proxy:$TAG .
docker push liftedkilt/foothillnorthchur.ch-proxy:$TAG