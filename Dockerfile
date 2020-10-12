FROM golang:1.15-alpine3.12

RUN apk --no-cache update && \
    apk add git protobuf

RUN go get -u github.com/golang/protobuf/protoc-gen-go

WORKDIR /app
COPY . .

EXPOSE 50051