# Build Stage
FROM golang:alpine3.20 AS builder

WORKDIR /build

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o /app ./cmd/serve

FROM migrate/migrate:4 AS migrator

# Final Stage
FROM alpine:3.20

EXPOSE 8001

COPY --from=builder /app /app
COPY --from=migrator /migrate /migrate
COPY --from=builder /build/db/migrations /migrations

CMD ["/app"]
