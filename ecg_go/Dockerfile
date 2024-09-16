FROM golang:1.23-alpine

WORKDIR /goSensor
COPY go.mod go.mod
COPY main.go main.go

RUN go build -o main .

CMD ["./main"]
