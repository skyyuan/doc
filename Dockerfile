FROM golang:1.9.2

COPY src/mindoc /go/src/mindoc

COPY src/mindoc/conf /go/bin/conf

COPY src/mindoc/static/fonts /go/bin/static/fonts

WORKDIR /go/src/mindoc

RUN go install mindoc

CMD ["/go/bin/mindoc"]
