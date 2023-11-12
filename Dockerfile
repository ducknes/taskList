FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod .
COPY ./ ./

RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]