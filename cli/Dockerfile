FROM golang:1.16.0-buster AS builder

RUN mkdir -p /go/src/github.com/seeis/cli

WORKDIR /go/src/github.com/seeis/cli

COPY constants ./constants
COPY engine/engine.go ./engine/
COPY client/client.go ./client/
COPY logger/logger.go ./logger/
COPY manager ./manager
COPY measurement/interface.go ./measurement/
COPY measurement/measurement_factory.go ./measurement/
COPY measurement/measurement.go ./measurement/
COPY sensors/cpu.go ./sensors/
COPY opts/opts.go ./opts/
COPY go.mod .
COPY go.sum .
COPY main.go .
RUN go build -o cli

FROM golang:1.16.0-buster
RUN mkdir -p /var/lib
WORKDIR /lib
COPY --from=builder /go/src/github.com/seeis/cli/cli .
CMD ./cli --unit=C --format=JSON --delta_duration=$CLI_DELTA_DURATION --total_duration=$CLI_TOTAL_DURATION --web_hook_url=$CLI_WEB_HOOK_URL