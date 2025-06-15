FROM golang:1.24.4-alpine AS builder 

WORKDIR /app

RUN apk add --no-cache git gcc g++ libc-dev

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o hospital-backend .

FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/hospital-backend .
COPY --from=builder /app/docs ./docs
COPY --from=builder /app/.env.example .env

# Set default port
EXPOSE 8080

# Run the binary
CMD ["./hospital-backend"]