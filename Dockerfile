FROM golang:1.8.1-alpine AS build-env

WORKDIR /go

RUN apk add --no-cache git

RUN go get -u github.com/golang/protobuf/proto && \
#    go get -u github.com/golang/protobuf/protoc-gen-go && \
    go get -u github.com/hooklift/httpclient && \
    go get -u github.com/Sirupsen/logrus && \
    go get -u golang.org/x/net/context && \
    go get -u google.golang.org/grpc && \
    go get -u google.golang.org/grpc/reflection

COPY . /go/src/github.com/avegao/iotArduino

RUN cd /go/src/github.com/avegao/iotArduino && go install


FROM alpine:3.5
WORKDIR /app
COPY --from=build-env /go/bin/iotArduino /app/iotArduino

EXPOSE 50000

CMD ./iotArduino
