FROM golang:1.16.5 AS development
# Add a work directory
WORKDIR /app
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy app files
COPY . .
# Install Reflex for development
RUN go install github.com/cespare/reflex@latest
# Expose port
EXPOSE 4000
# Start app
CMD reflex -g '*.go' go run server.go --start-service

FROM golang:1.16.5 AS builder
# Define build env
ENV GOOS linux
ENV CGO_ENABLED 0
# Add a work directory
WORKDIR /app
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy app files
COPY . .
# Build app
RUN go build -o app

FROM alpine:3.14 AS production
# Add certificates
RUN apk add --no-cache ca-certificates
# Copy built binary from builder
COPY --from=builder app .
# Expose port
EXPOSE 4000
# Exec built binary
CMD ./app