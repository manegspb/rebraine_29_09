FROM golang:1.17 as builder

WORKDIR /go/src/httpserver

COPY main.go ./

RUN GO111MODULE=off \
    CGO_ENABLED=0 \
    GOOS=linux \
    go build  -o httpserver main.go

FROM alpine:3.11

COPY --from=builder /go/src/httpserver/httpserver /usr/bin/httpserver

EXPOSE 8080
ENTRYPOINT ["/usr/bin/httpserver"]