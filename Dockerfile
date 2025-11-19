FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/api

FROM alpine:latest
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/app /usr/local/bin/app

EXPOSE 8080

CMD ["app"]