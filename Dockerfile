FROM golang:1.23-alpine

WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary
RUN go build -o /dx-agent-dummy

# Use a minimal image for the final executable (optional)
FROM alpine:latest
COPY --from=0 /dx-agent-dummy /dx-agent-dummy

ENTRYPOINT ["/dx-agent-dummy"]
