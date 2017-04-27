FROM golang:alpine
COPY demo.go /go/
RUN apk update && apk add git && \
    go get github.com/garyburd/redigo/redis && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
