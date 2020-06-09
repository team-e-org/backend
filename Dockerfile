FROM golang:1.14.2-alpine
COPY . /go/src/app
WORKDIR /go/src/app
RUN apk update && apk upgrade && apk add --no-cache git \
    && go get -u \
    && apk del --purge git \
    && go get github.com/cosmtrek/air

ENV SERVER_PORT=3000

CMD air -d
