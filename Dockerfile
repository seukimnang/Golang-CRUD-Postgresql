# Build stage
FROM golang:latest AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the pre-built binary from the previous stage
COPY --from=builder /app/main .

# Expose port 80 to the outside world
EXPOSE 8888

# Command to run the executable
CMD ["./main"]
