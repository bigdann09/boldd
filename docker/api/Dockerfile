FROM golang:1.24 AS builder
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o boldd cmd/api/main.go

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /ecommerce-api .
# COPY migrations ./migrations
EXPOSE ${APP_PORT}
CMD ["/app/boldd"]