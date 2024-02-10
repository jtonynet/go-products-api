# Primeira etapa: construção
FROM golang:1.21.1-alpine AS api

WORKDIR /usr/src/app

COPY .env  ./
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

# TODO: Fix to compile and run binary in another Dockerfile step
# RUN go build -o main ./cmd/api/main.go
# CMD["./main"]

CMD ["go", "run", "cmd/api/main.go"]

