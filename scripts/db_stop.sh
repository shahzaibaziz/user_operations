#!/bin/bash

if [ "$(docker ps -a -q -f name=mongodb_dev)" ]; then
  docker rm -f mongodb_dev
fi
