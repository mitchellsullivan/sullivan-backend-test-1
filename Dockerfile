# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /sullivan-backend-test-1 .

EXPOSE 8099

CMD [ "/sullivan-backend-test-1" ]