#!/usr/bin/env bash

set -e
[[ -z $DEBUG ]] || set -x

services="rello users model"

for service in $services; do
	(
		echo "building $service"
		cd $service;
		mkdir -p target
		go build -o ./target/newrun ./${service}-service/cmd/${service}-server/
		mv ./target/newrun ./target/run
	)
done
