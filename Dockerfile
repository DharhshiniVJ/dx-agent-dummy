# Use Go 1.22 base image
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Cache go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN go build -o dx-agent-dummy ./...

# Final minimal image
FROM alpine:latest
COPY --from=builder /app/dx-agent-dummy /usr/local/bin/dx-agent-dummy

ENTRYPOINT ["dx-agent-dummy"]
