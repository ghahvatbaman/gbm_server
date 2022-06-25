FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o ghahvatbaman ./cmd/api

RUN chmod -x /app/ghahvatbaman

FROM alpine:latest

RUN mkdir /app

COPY  --from=builder /app/ghahvatbaman /app

CMD [ "/app/ghahvatbaman" ]