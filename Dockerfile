FROM golang:1.24 AS builder

WORKDIR /app

# Copy all files
COPY . .

RUN go mod tidy

RUN go build -o migrate ./cmd/migrate

RUN go build -o app ./cmd/app

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/migrate /app/migrate
COPY --from=builder /app/app /app/app

EXPOSE 8080

CMD ["/app/app"]
