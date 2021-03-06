FROM golang:1.8.1-alpine AS build-env

RUN mkdir -p /go/src/github.com/avegao/iot-arduino

WORKDIR /go/src/github.com/avegao/iot-arduino

RUN apk add --no-cache git glide

COPY glide.yaml glide.yaml
COPY glide.lock glide.lock

RUN glide install

COPY . .

RUN go install

########################################################################################################################

FROM alpine:3.5
WORKDIR /app
COPY --from=build-env /go/bin/iot-arduino /app/iot-arduino

EXPOSE 50000

LABEL maintainer="Álvaro de la Vega Olmedilla <alvarodlvo@gmail.com>"

ENTRYPOINT ["./iot-arduino"]
