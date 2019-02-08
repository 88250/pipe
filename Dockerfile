FROM alpine:3.7
LABEL maintainer = "abcdsxg@gmail.com"

ENV PIPE_VERSION v1.8.6
ENV GLIBC_VERSION 2.27-r0

WORKDIR /opt/

RUN set -ex && \
    apk --no-cache add ca-certificates && \
    wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub && \
    wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-${GLIBC_VERSION}.apk && \
    apk add glibc-${GLIBC_VERSION}.apk && \
    wget -O pipe${PIPE_VERSION}.zip https://github.com/b3log/pipe/releases/download/${PIPE_VERSION}/pipe-${PIPE_VERSION}-linux.zip && \
    unzip pipe${PIPE_VERSION}.zip && \
    chmod +x pipe && \
    rm -f pipe${PIPE_VERSION}.zip glibc-${GLIBC_VERSION}.apk

CMD ["/opt/pipe"]