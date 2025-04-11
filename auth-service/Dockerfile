# Dockerfile
FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o auth-service ./cmd/main.go

EXPOSE 50051

CMD ["./auth-service"]
