# Build stage
FROM debian:bullseye-slim AS builder

# Install bazel and required dependencies
RUN apt-get update && apt-get install -y \
    curl \
    gnupg \
    && curl -fsSL https://bazel.build/bazel-release.pub.gpg | gpg --dearmor > /etc/apt/trusted.gpg.d/bazel.gpg \
    && echo "deb [arch=amd64] https://storage.googleapis.com/bazel-apt stable jdk1.8" | tee /etc/apt/sources.list.d/bazel.list \
    && apt-get update && apt-get install -y bazel \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /build

# Copy the source code
COPY . .

# Build the binary using bazel
RUN bazel build //cmd/arithmetic_server:server

# Runtime stage
FROM debian:bullseye-slim

WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /build/bazel-bin/cmd/arithmetic_server/server /app/server

# Expose gRPC port
EXPOSE 50001

# Run the server
CMD ["/app/server"]
