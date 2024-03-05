From golang:1.21.5

WORKDIR /go/src/app

COPY . .

RUN go build -o main main.go

CMD ["./main"]