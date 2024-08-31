FROM golang:1.22.5-alpine3.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY src src

RUN go build src/main.go

EXPOSE 8080

ENTRYPOINT [ "main" ] 
