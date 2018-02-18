FROM golang:1.9.4-alpine3.6
RUN apk add --update go git
RUN go get github.com/gorilla/mux github.com/lib/pq
RUN mkdir -p /go/src/app
WORKDIR /go/src/app

ADD . /go/src/app
COPY ./library-api /bin/library-api

EXPOSE 8080

COPY ./wait-for-it.sh /bin/wait-for-it.sh
RUN chmod +x /bin/wait-for-it.sh