FROM golang:1.23.2 AS builder

WORKDIR /app

COPY go.mod go.sum  ./
RUN go mod download

COPY . .

RUN go build -o /app/serve ./cmd/judgino/main.go
RUN go build -o /app/code-runner ./cmd/code-runner/main.go
RUN go build -o /app/create-admin ./cmd/create-admin/main.go

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/serve /usr/local/bin/
COPY --from=builder /app/code-runner /usr/local/bin/
COPY --from=builder /app/create-admin /usr/local/bin/
COPY assets ./assets
COPY templates ./templates
COPY static ./static

RUN mkdir -p /app/uploads


