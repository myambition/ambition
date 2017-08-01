#!/usr/bin/env bash

set -e
[[ -z $DEBUG ]] || set -x

cur_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

services="rello users model"

for service in $services; do
	(
		echo "building $service"
		cd $cur_dir/$service;
		mkdir -p target
		go build -o ./target/newrun ./${service}-service/cmd/${service}_d/
		mv ./target/newrun ./target/run
	)
done
