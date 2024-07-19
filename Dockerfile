FROM golang:1.22.5
WORKDIR /app

COPY ./Server/go.mod ./Server/go.sum ./
RUN go mod download
COPY ./Server ./Server
COPY ./Client /app/Client

RUN go build -o /app/Server/main ./Server/main.go

WORKDIR /app/Server
EXPOSE 8080

CMD ["./main"]
