FROM golang:latest

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .
RUN go build -o notes_service ./cmd/main/main.go

EXPOSE 8088

CMD ["./notes_service"]