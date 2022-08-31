FROM golang:1.18.0

WORKDIR /go/src/memorystore-sample

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o main .

CMD ["./main"]
