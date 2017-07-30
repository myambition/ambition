#!/usr/bin/env bash

set -e
[[ -z $DEBUG ]] || set -x

epoc="$(date -u "+%s")"
tag="myambition/ambition-mysql:$epoc"
echo $tag

docker build -t "$tag"  .
[[ -n $NO_PUSH ]] || docker push "$tag"
