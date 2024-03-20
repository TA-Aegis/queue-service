FROM golang:1.22-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go mod download

RUN go build -o binary application/*.go

EXPOSE 8080

EXPOSE 9090

ENTRYPOINT ["/app/binary"]