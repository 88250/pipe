FROM node:alpine as NODE_BUILD
ADD . /tmp
RUN cd /tmp/console && npm install && npm run build && cd ../theme && npm install && npm run build

FROM golang:alpine as GOLANG_BUILD
WORKDIR /go/src/github.com/b3log/pipe/
COPY --from=NODE_BUILD /tmp .
RUN go build -i -v

FROM apline
LABEL maintainer="Liang Ding<d@b3log.org>"

WORKDIR /opt/pipe
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/${TZ} /etc/localtime && echo ${TZ} > /etc/timezone \
    && cp /go/src/github.com/b3log/pipe/* /opt/pipe/ \
    && rm -rf /tmp/*

EXPOSE 5897

ENTRYPOINT [ "pipe", "--version"]
