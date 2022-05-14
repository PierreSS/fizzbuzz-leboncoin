FROM golang:buster

COPY . /app
WORKDIR /app

RUN export GOPATH=$HOME/go
RUN export PATH=/usr/local/go/bin:$PATH:$GOPATH/bin
RUN go get -d
RUN go mod download
RUN go generate
RUN make

ENTRYPOINT [ "./bin/fizzbuzz-leboncoin" ]
