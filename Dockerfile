FROM golang:1.8.1-alpine

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

EXPOSE 50000

CMD ./bin/iotArduino
