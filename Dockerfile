FROM golang:1.22.5-alpine3.20

WORKDIR /app 

COPY . .

RUN go build -o main main.go 

EXPOSE 8080

ENV path="." \
    name="dev" \
    ext="env"

CMD ["/app/main"]
