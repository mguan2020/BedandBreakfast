FROM golang:latest

WORKDIR /app

COPY . /app

RUN go build -o bookings cmd/web/main.go cmd/web/middleware.go cmd/web/routes.go

EXPOSE 8080

ENTRYPOINT ["./run.sh"]
