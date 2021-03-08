FROM golang

ADD . /go/src/github.com/alelaca/booking-backend

RUN go install github.com/alelaca/booking-backend/cmd/server@latest

ENTRYPOINT /go/bin/booking-backend

EXPOSE 9000