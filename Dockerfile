FROM alpine:3.3

MAINTAINER Oleg Anedchenko <oanedchenko@gmail.com>
LABEL "app"="autonet"

RUN apk update && apk add curl && rm -rf /var/cache/apk/*

ENV DB_URL http://arango:8529
ENV DB_NAME autonet
ENV PORT 9090
ENV HOST 0.0.0.0
ENV APP /auto1
WORKDIR $APP
ADD autonet $APP/autonet
EXPOSE $PORT

ENTRYPOINT $APP/autonet
