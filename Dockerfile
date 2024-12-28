FROM golang:1.23.2-alpine

WORKDIR /app

COPY . .

RUN go build -o server ./cmd/server

EXPOSE 8000

CMD ["./server"]
