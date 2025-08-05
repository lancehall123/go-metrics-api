FROM golang:1.21-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o go-metrics-api

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/go-metrics-api .
EXPOSE 8080
CMD ["./go-metrics-api"]
