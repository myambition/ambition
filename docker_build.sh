#!/usr/bin/env bash

set -e
[[ -z $DEBUG ]] || set -x

docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD

cur_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

services="rello users model"

for service in $services; do
	(
		echo "building $service"
		tag="myambition/ambition-${service}:latest"
		echo $tag
		docker build -t $tag --build-arg service=${service} .
		docker push $tag &
	)
done
