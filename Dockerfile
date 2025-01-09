# Build Stage
FROM golang:alpine3.20 AS builder

WORKDIR /build

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o /app ./cmd/serve

# Final Stage
FROM alpine:3.20

EXPOSE 8001

COPY --from=builder /app /app
CMD ["/app"]
