FROM golang:1.27-trixie

RUN apt-get update && \
    apt-get install -y curl gnupg2 \
                       git

RUN ln -s /usr/local/go/bin/go /usr/local/bin/ && \
    ln -s /usr/local/go/bin/gofmt /usr/local/bin/

RUN adduser app

ENV GQL_VERSION="0.17.94"
RUN cd /root && \
    wget https://github.com/99designs/gqlgen/archive/refs/tags/v${GQL_VERSION}.tar.gz -O gqlgen.tar.gz && \
    tar -xf gqlgen.tar.gz && \
    cd gqlgen-${GQL_VERSION} && \
    go build main.go && \
    mv main /bin/gqlgen

WORKDIR /app
