FROM golang:latest

WORKDIR /app

COPY ./go-notes-service .

RUN go mod download

RUN go build -o main ./cmd

EXPOSE 8080

CMD ["./main"]
