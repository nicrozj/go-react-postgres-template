FROM golang:1.24.1-alpine AS dev
WORKDIR /app

RUN go install github.com/air-verse/air@latest
COPY go.mod go.sum ./
RUN go mod download
CMD ["air"]