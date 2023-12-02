FROM golang:1.19-buster

WORKDIR /app

COPY . /app

RUN go mod download

EXPOSE 8080

CMD ["go", "run", "/app/cmd/bot_api/main.go"]