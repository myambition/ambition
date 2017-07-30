#!/usr/bin/env bash

set -e
[[ -z $DEBUG ]] || set -x

epoc="$(date -u "+%s")"
tag="myambition/ambition-mysql:$epoc"
latest="myambition/ambition-mysql:latest"
echo $tag

docker build -t "$tag" -t $latest .
[[ -n $NO_PUSH ]] || docker push "$tag" && docker push $latest
