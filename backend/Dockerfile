FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o brokerx .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/brokerx .

EXPOSE 6969

CMD [ "./brokerx" ]