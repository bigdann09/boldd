FROM golang:1.24 AS builder
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o boldd cmd/api/main.go

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/boldd /app/boldd
EXPOSE ${APP_PORT:-8003}
CMD ["/app/boldd"]