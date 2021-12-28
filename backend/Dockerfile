FROM golang:1.16.0-buster AS builder

RUN mkdir -p /go/src/github.com/I544091/backend

WORKDIR /go/src/github.com/I544091/backend

COPY constants ./constants
COPY db ./db
COPY server ./server
COPY logger ./logger
COPY manager ./manager
COPY main.go .
COPY go.mod .
COPY go.sum .
RUN go build -o backend

FROM golang:1.16.0-buster

RUN mkdir -p /var/lib

WORKDIR /lib

COPY --from=builder /go/src/github.com/I544091/backend/backend .

EXPOSE 8000

CMD ["./backend"]