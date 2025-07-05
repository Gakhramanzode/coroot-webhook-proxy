
FROM golang:1.23-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=direct

WORKDIR /build

RUN apk add --no-cache git

COPY go.sum .

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:3.17

COPY --from=builder /build/main /

COPY config /config

USER nobody

ENTRYPOINT ["/main"]
