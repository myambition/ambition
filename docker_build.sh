#!/usr/bin/env bash

set -e
[[ -z $DEBUG ]] || set -x

docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD

cur_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

services="rello users model"
tag_postfix="latest"

if [[ "$CI" == "true" ]]; then
	curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -;
	sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable";
	sudo apt-get update;
	sudo apt-get -y install docker-ce;
	tag_postfix="prod-latest"
fi

for service in $services; do
	(
		echo "building $service"
		tag="myambition/${service}:${tag_postfix}"
		echo $tag
		docker build -t $tag --build-arg service=${service} .
		docker push $tag
	)
done
