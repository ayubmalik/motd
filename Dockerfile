FROM golang:1.11.5-alpine3.7

COPY . /go/src/github.com/ayubmalik/motd

RUN go install /go/src/github.com/ayubmalik/motd

ENTRYPOINT /go/bin/motd

EXPOSE 8000
