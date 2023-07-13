FROM golang:1.16-alpine

WORKDIR /goSensor
COPY go.mod go.mod
COPY goSensor.go goSensor.go

RUN go build -o main .

CMD ["./main"]
