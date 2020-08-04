FROM golang:1.14

RUN apt-get update || exit 0
RUN apt-get upgrade -y
RUN mkdir -p /go/src/github.com/request-limit
COPY . /go/src/github.com/request-limit
WORKDIR /go/src/github.com/request-limit
RUN  go mod tidy
RUN go test -cover ./...
RUN go install /go/src/github.com/request-limit