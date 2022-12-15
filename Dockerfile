FROM golang:1.18

WORKDIR /go/src/app

COPY main.go .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
