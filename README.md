# Go Image Optimizer Microservice 🚀

A high-performance microservice for image processing, designed to be resilient and lightweight.
This project demonstrates Go's efficiency for CPU-bound tasks compared to Node.js.

![Go](https://img.shields.io/badge/Go-1.25.4-00ADD8?style=flat&logo=go)
![Docker](https://img.shields.io/badge/Docker-Alpine-2496ED?style=flat&logo=docker)
![Size](https://img.shields.io/badge/Image_Size-~18MB-success)

## ⚡ Key Features

- **Native Image Processing**: Decoding, resizing (Lanczos), and optimized JPEG encoding without heavy system dependencies.
- **Memory Protection (Semaphore Pattern)**: Uses buffered Go Channels to limit concurrency and prevent OOM (Out Of Memory) kills under heavy load.
- **Fail-Fast Architecture**: Immediate rejection of excess requests (Status 503) to maintain low latency for active users.
- **Graceful Shutdown**: OS signal handling (SIGTERM/SIGINT) to complete in-flight requests before container shutdown (Zero data loss).
- **Docker Multi-Stage Build**: Static production binary weighing less than 20 MB.

## 🛠️ Installation & Setup

### Via Docker (Recommended)

```bash
# Build the lightweight image
docker build -t go-optimizer .

# Run the container (Port 8080)
docker run -p 8080:8080 --name optimizer go-optimizer
```

### Local Development

```bash
go mod download
go run .
```

## 🧪 Benchmark & Performance

Test performed with **k6** (50 VUs, 2MB image uploads, 30s):

| Metric | Result |
|--------|--------|
| Throughput | ~50 req/sec (CPU-bound) |
| Max Memory | Stable at ~250MB (thanks to semaphore) |
| p95 Latency | < 1.5s |
| Docker Image | 20.5 MB |

## 📐 Code Architecture

- `main.go`: Entry point, server configuration and Graceful Shutdown handling.
- `handle_upload.go`: Business logic and concurrency management (Semaphore).
- **Goroutines**: Each request is handled in its own lightweight thread.

## 🔗 Author

Antoine Rospars - [Portfolio](https://antoinerospars.dev) - [GitHub](https://github.com/P4ST4S)
