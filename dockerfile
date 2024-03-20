FROM golang:1.22-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go mod download

RUN go build -o binary application/rest/*.go

EXPOSE 8080

ENTRYPOINT ["/app/binary"]