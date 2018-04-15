#!/bin/sh

APP=autonet

echo "removing old ${APP} container..."
docker rm -f $APP

echo "loading new ${APP} image from file..."
gunzip -c /root/deploy/auto1/${APP}.tgz | docker load
if [ $? -eq 0 ]; then
    echo "image loaded successfully"
else
    echo "fail to load image"
    exit 1
fi

echo "starting ${APP}..."
docker run -d \
        --name $APP \
        -e DB_URL=tcp://172.17.0.1:8529 \
        -e DB_PASSWORD=123 \
        -e DB_USER=root \
        auto1/${APP}:latest
