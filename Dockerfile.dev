FROM golang:1.21.4-alpine3.18

WORKDIR /app

COPY . /app

RUN go mod download
RUN apk add --update --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    echo "Asia/Tokyo" > /etc/timezone && \
    apk del tzdata

EXPOSE 8080

CMD ["go", "run", "/app/cmd/bot_api/main.go"]