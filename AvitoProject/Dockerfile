FROM golang:alpine

RUN apk add git
WORKDIR /go/src/AvitoProject/
COPY . .
WORKDIR /go/src/AvitoProject/cmd
RUN go get -d -v ./...
RUN go build
CMD ./cmd
