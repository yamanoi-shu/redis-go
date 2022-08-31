FROM golang:1.17

WORKDIR /go/src/redis-go-sample

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

RUN GOOS=linux GOARCH=amd64 go build -o main main.go

CMD ["./main"]
