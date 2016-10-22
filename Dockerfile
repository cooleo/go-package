FROM golang:onbuild
ADD . /go/src/app/greet
EXPOSE 8080
