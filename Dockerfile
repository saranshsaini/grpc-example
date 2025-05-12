# Build stage
FROM debian:bullseye-slim AS builder

# Install required dependencies
RUN apt-get update && apt-get install -y \
    curl \
    wget \
    git \
    gcc \
    g++ \
    build-essential \
    && rm -rf /var/lib/apt/lists/*

# Install Go and add to PATH
RUN wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz \
    && rm go1.21.5.linux-amd64.tar.gz

# Add Go and its bin directory to PATH
ENV PATH="/usr/local/go/bin:/root/go/bin:${PATH}"

# Install Bazelisk
RUN go install github.com/bazelbuild/bazelisk@latest

# Set Bazel version
RUN echo "USE_BAZEL_VERSION=7.1.1" > /root/.bazeliskrc

WORKDIR /build

# Copy the source code
COPY . .

# Build the binary using bazel and find its location
RUN bazelisk build //cmd/arithmetic_server:server && \
    cp bazel-bin/cmd/arithmetic_server/server_/server /build/server

# Runtime stage
FROM debian:bullseye-slim

WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /build/server /app/server

# Expose gRPC port
EXPOSE 50051

# Run the server
CMD ["/app/server"]
