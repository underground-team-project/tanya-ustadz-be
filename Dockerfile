FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /home/tanyaustadz

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o /home/tanyaustadz/bin/app

FROM alpine:latest

WORKDIR /home/tanyaustadz

COPY --from=builder /home/tanyaustadz/bin/app /home/tanyaustadz/bin/

ENTRYPOINT ["/home/tanyaustadz/bin/app"]
