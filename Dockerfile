FROM golang

ADD . /go/src/github.com/alelaca/booking-backend

RUN go install github.com/alelaca/booking-backend@latest

ENTRYPOINT /go/bin/booking-backend

EXPOSE 9000