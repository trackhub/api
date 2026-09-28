FROM golang:1.27-trixie

RUN apt-get update && \
    apt-get install -y curl gnupg2 \
                       git

RUN ln -s /usr/local/go/bin/go /usr/local/bin/ && \
    ln -s /usr/local/go/bin/gofmt /usr/local/bin/

RUN adduser app

USER app

ENV PATH="usr/local/go/bin/:${PATH}"

WORKDIR /app
