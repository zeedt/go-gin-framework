FROM golang:latest

MAINTAINER Saheed Yusuf

ENV GIN_MODE=release
ENV PORT=3003


RUN cd /go/src

RUN mkdir myapp

WORKDIR /go/src/myapp

COPY public /go/src/myapp/public
COPY src /go/src/myapp/src
COPY templates /go/src/myapp/templates

RUN cd /go/src/myapp

#RUN go get -u github.com/gin-gonic/gin

COPY dependencies /go/src

RUN cd /go/src

RUN go build myapp/src/app

RUN pwd

EXPOSE $PORT

ENTRYPOINT ["./app"]
