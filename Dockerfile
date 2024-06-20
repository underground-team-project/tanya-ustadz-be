ARG GOLANG=golang:alpine

FROM ${GOLANG} AS base

FROM ${GOLANG} AS builder

WORKDIR /src
COPY ./ ./

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -o tanyaustadz ./ 

FROM alpine AS alpine

FROM scratch 

COPY --from=builder /src/tanyaustadz /usr/local/bin/tanyaustadz

ENTRYPOINT ["/usr/local/bin/tanyaustadz"]
