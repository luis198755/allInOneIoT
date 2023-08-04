FROM golang:latest

WORKDIR /goCliInflux
COPY go.sum go.sum
COPY go.mod go.mod
COPY goCli.go goCli.go

RUN go mod tidy

RUN go build -o main .

CMD ["./main"]
