FROM golang:1.25.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o my-optimizer .


FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/my-optimizer .

EXPOSE 8080

CMD ["./my-optimizer"]