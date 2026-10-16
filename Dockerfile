FROM node:24 AS node_build
WORKDIR /go/src/github.com/88250/pipe/
ADD . /go/src/github.com/88250/pipe/
# 为 Nuxt 2 使用的 Webpack 4 提供 OpenSSL 兼容算法。
ENV NODE_OPTIONS=--openssl-legacy-provider
RUN cd console && npm install && npm run build && cd ../theme && npm install && npm run build && \
    rm -rf node_modules && cd ../console && rm -rf node_modules

FROM golang:alpine AS go_build
WORKDIR /go/src/github.com/88250/pipe/
COPY --from=node_build /go/src/github.com/88250/pipe/ /go/src/github.com/88250/pipe/
ENV GO111MODULE=on
RUN apk add --no-cache gcc musl-dev git && go build -v

FROM alpine:latest
LABEL maintainer="Liang Ding<845765@qq.com>"
WORKDIR /opt/pipe/
COPY --from=go_build /go/src/github.com/88250/pipe/ /opt/pipe/
RUN apk add --no-cache ca-certificates tzdata

ENV TZ=Asia/Shanghai
EXPOSE 5897

ENTRYPOINT [ "/opt/pipe/pipe" ]
