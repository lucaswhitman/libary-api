FROM golang:1.8

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

ADD . /go/src/app
COPY ./library-api /bin/library-api

EXPOSE 8080

COPY ./wait-for-it.sh /bin/wait-for-it.sh
RUN chmod +x /bin/wait-for-it.sh
CMD ["/bin/wait-for-it.sh", APP_ENV]
RUN go get github.com/gorilla/mux github.com/lib/pq