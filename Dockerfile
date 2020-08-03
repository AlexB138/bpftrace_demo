FROM golang:1.14.6-stretch

COPY demo.go .

RUN go build -gcflags '-l' -o demo ./demo.go
