FROM golang:1.19-bullseye

RUN apt-get update && \
    apt-get install -y curl gnupg2 \
                       git

ENV GQL_VERSION="0.17.20"

#RUN cd /root && \
#    wget https://github.com/99designs/gqlgen/archive/refs/tags/v${GQL_VERSION}.tar.gz -O gqlgen.tar.gz && \
#    tar -xf gqlgen.tar.gz && \
#    cd gqlgen-${GQL_VERSION} && \
#    go build main.go && \
#    mv main /bin/gqlgen

WORKDIR /app
