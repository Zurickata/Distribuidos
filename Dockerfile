FROM golang:latest

WORKDIR /app
COPY servidor/servidor.go .

RUN go build servidor.go

EXPOSE 8080

CMD ["./servidor"]
