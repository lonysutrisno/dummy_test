FROM golang:latest

RUN mkdir -p /go/src/github.com/lonysutrisno/dummy_test
ADD . /go/src/github.com/lonysutrisno/dummy_test
WORKDIR /go/src/github.com/lonysutrisno/dummy_test

EXPOSE 8000

CMD ["./dummy_test"]
