FROM golang:1.23.1-alpine AS builder

RUN apk add --no-cache \
    git

RUN --mount=type=cache,target=/go/pkg/mod \
    mkdir /benchstat-git && cd /benchstat-git && \
    git clone https://go.googlesource.com/perf .&& \
    go build -o ./out/benchstat ./cmd/benchstat

FROM golang:1.23.1-alpine AS benchstat

COPY --from=builder /benchstat-git/out/benchstat /usr/local/bin/benchstat

ENTRYPOINT [ "/usr/local/bin/benchstat" ]
