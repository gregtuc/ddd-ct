# Use the official Go image to build the application.
FROM golang:1.21.5 as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ddd-ct ./cmd/webserver

FROM debian:buster-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/ddd-ct .

EXPOSE 8080

CMD ["./ddd-ct"]
