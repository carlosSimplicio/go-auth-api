FROM golang:1.22.5-alpine3.20 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY src src

RUN go build src/main.go

FROM scratch

WORKDIR /

EXPOSE 8080

ENV MYSQL_PASSWORD 953042
ENV MYSQL_USER root
ENV MYSQL_ADDRESS 172.19.0.2:3306

COPY --from=builder /app .

ENTRYPOINT [ "./main" ] 
