#!/usr/bin/env bash

./services/build.sh
docker-compose restart model users rello
