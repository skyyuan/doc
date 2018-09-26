FROM golang:1.9.2

COPY src/mindoc /go/src/mindoc

WORKDIR /go/src/mindoc

RUN go build

CMD ["/go/src/mindoc/mindoc"]
