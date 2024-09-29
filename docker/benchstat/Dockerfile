FROM golang:1.23.1-alpine AS build

RUN go install golang.org/x/perf/cmd/benchstat@latest

ENTRYPOINT [ "benchstat" ]
