FROM golang:onbuild
ADD . /go/src/greet
EXPOSE 8080
