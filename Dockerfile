FROM golang:1.17-buster

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

CMD ["go", "run", "./cmd/server/main.go"]
