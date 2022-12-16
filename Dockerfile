ARG GO_VERSION=1.18

FROM golang:${GO_VERSION} AS builder

WORKDIR /

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN go build -o ./main main.go

EXPOSE 8080

ENTRYPOINT ["./main"]
