FROM golang:latest AS builder
WORKDIR /go/src/app

COPY . .

ENV GO111MODULE=on
RUN go mod download

RUN go install -v ./...
RUN go build .

EXPOSE 90
CMD ["go_microservice_template", "-b", ":90"]
