#!/usr/bin/env bash

set -e
[[ -z $DEBUG ]] || set -x

services="rello users model"

for service in $services; do
	(
		cd $service;
		truss *.proto;
	)
done
