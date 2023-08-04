FROM golang:latest

WORKDIR /mockSensorGo
COPY go.mod go.mod
COPY main.go main.go

RUN go mod tidy

RUN go build -o main .

CMD ["./main"]
