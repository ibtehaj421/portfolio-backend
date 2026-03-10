# Build Stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app. CGO_ENABLED=0 ensures a static binary.
RUN CGO_ENABLED=0 GOOS=linux go build -o /api-server ./cmd/api/main.go

# Final Stage
FROM alpine:latest

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /api-server .

# Expose the port
EXPOSE 8080

# Command to run the executable
CMD ["./api-server"]