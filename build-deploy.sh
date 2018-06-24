#!/bin/sh

APP=autonet
IMAGE=auto1/autonet:latest
P=`pwd`

cd cmd/autonet-api-documentation-server
CC=$(which musl-gcc) go build -v -o $P/$APP --ldflags '-w -linkmode external -extldflags "-static"'

cd $P

docker build -t $IMAGE .
docker save $IMAGE | gzip > ${APP}.tgz
scp ${APP}.tgz root@10f20.k.time4vps.cloud:/root/deploy/auto1
ssh root@10f20.k.time4vps.cloud /root/deploy/auto1/redeploy-autonet.sh

rm autonet autonet.tgz

