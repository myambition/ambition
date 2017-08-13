#!/usr/bin/env bash

set -e
[[ -z $DEBUG ]] || set -x

cur_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
repo_root=$(git rev-parse --show-toplevel 2> /dev/null)

docker_compose_dir="${repo_root}"

(
	cd $docker_compose_dir
	docker-compose -f ${docker_compose_dir}/docker-compose.release.yml config > ${cur_dir}/docker-compose.yml
)
