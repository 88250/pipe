#!/usr/bin/env bash


PIPE_VERSION="v1.0"
PIPE_CONTAINER_NAME="Marveliu-blog"

docker run --rm -v $HOME/gopath:/go marveliu/go1.11-alpine-gcc:v1.0 sh -c 'cd /go/src/github.com/b3log/pipe && CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -i -v'

docker stop ${PIPE_CONTAINER_NAME} | docker rm ${PIPE_CONTAINER_NAME}
docker rmi marveliu/marveliu-pipe:${PIPE_VERSION}
#
docker build . --tag marveliu/marveliu-pipe:${PIPE_VERSION}
docker run --rm -p 5897:5897 --name ${PIPE_CONTAINER_NAME} marveliu/marveliu-pipe:${PIPE_VERSION} sh -c 'sh ./pipe'
