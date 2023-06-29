FROM golang:1.17

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

ENV GOFLAGS=-mod=vendor

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
