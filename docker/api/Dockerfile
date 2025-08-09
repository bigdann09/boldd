FROM golang:1.24 AS builder
WORKDIR /src
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o boldd cmd/api/main.go

# Final stage
FROM alpine:latest
RUN mkdir -p /app/boldd
COPY --from=builder /src/boldd /app/boldd/boldd
WORKDIR /app/boldd
EXPOSE ${APP_PORT:-8003}
CMD ["./boldd"]