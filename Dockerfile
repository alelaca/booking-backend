FROM golang

ADD . /go/src/github.com/alelaca/booking-backend

RUN go install github.com/alelaca/booking-backend/cmd/server@master

ENTRYPOINT /go/bin/server

EXPOSE 9000