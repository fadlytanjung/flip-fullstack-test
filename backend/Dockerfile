# Backend Dockerfile for Go
# Multi-stage build for optimized production image
# Minimal final image size with Alpine Linux

# Stage 1: Builder
FROM golang:1.23-alpine AS builder

# Install required build tools including gcc for CGO/SQLite
RUN apk add --no-cache git ca-certificates tzdata gcc musl-dev

WORKDIR /build

# Copy go mod files first (for better Docker layer caching)
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application with CGO enabled for SQLite
# -ldflags="-w -s" to reduce binary size and remove debugging info
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o server \
    ./cmd/server

# Stage 2: Final Runtime Image
FROM alpine:latest

# Install only essential runtime dependencies including libgcc for CGO
RUN apk --no-cache add ca-certificates tzdata libgcc

# Create non-root user for security
RUN addgroup -g 1001 -S appuser && \
    adduser -u 1001 -S appuser -G appuser

WORKDIR /app

# Copy binary from builder
COPY --from=builder --chown=appuser:appuser /build/server .

# Set environment variables
ENV PORT=9000
ENV TZ=UTC

# Expose port
EXPOSE 9000

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:9000/api/health || exit 1

# Switch to non-root user
USER appuser

# Start the application
CMD ["./server"]

