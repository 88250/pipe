FROM node:10.15.3 as NODE_BUILD
WORKDIR /go/src/github.com/88250/pipe/
ADD . /go/src/github.com/88250/pipe/
RUN cd console && npm install && npm run build && cd ../theme && npm install && npm run build && \
    rm -rf node_modules && cd ../console && rm -rf node_modules

FROM golang:alpine as GO_BUILD
WORKDIR /go/src/github.com/88250/pipe/
COPY --from=NODE_BUILD /go/src/github.com/88250/pipe/ /go/src/github.com/88250/pipe/
ENV GO111MODULE=on
RUN apk add --no-cache gcc musl-dev git && go build -i -v

FROM alpine:latest
LABEL maintainer="Liang Ding<d@b3log.org>"
WORKDIR /opt/pipe/
COPY --from=GO_BUILD /go/src/github.com/88250/pipe/ /opt/pipe/
RUN apk add --no-cache ca-certificates tzdata

ENV TZ=Asia/Shanghai
EXPOSE 5897

ENTRYPOINT [ "/opt/pipe/pipe" ]
