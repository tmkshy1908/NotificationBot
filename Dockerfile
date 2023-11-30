FROM golang:1.19-buster

RUN mkdir app

WORKDIR /app

COPY . /app

EXPOSE 8080

CMD ["go", "run", "/app/cmd/bot_api/main.go"]