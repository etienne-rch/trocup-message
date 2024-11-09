# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Copy only the dependency files first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application with additional flags
RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Final stage
FROM alpine:latest

WORKDIR /app

# Add ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates  

# Copy the binary from builder
COPY --from=builder /app/app .

# Create a non-root user
RUN adduser -D appuser
USER appuser

EXPOSE 5004

# Add healthcheck
HEALTHCHECK --interval=30s --timeout=3s \
  CMD wget --no-verbose --tries=1 --spider http://localhost:5004/health || exit 1

CMD ["./app"]