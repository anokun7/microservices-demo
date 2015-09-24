FROM golang:onbuild

RUN go get github.com/garyburd/redigo/redis
