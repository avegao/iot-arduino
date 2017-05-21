package main

import (
    "github.com/Sirupsen/logrus"
    "net"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    pb "github.com/avegao/iot-arduino/proto"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "time"
    "os"
    "github.com/o1egl/gormrus"
)

const (
    PORT = ":50000"
    ENV_DEV = "dev"
    ENV_TEST = "test"
    ENV_PRE = "pre"
    ENV_PROD = "prod"
)

var database *gorm.DB
var ENV = os.Getenv("GO_ENV")

func isDev() bool {
    return ENV_DEV == ENV
}

func isTest() bool {
    return ENV_TEST == ENV
}

func isPre() bool {
    return ENV_PRE == ENV
}

func isProd() bool {
    return ENV_PROD == ENV
}

func initLogger() {
    logLevel := logrus.ErrorLevel

    if isDev() || isTest() {
        logLevel = logrus.DebugLevel
    } else if isPre() {
        logLevel = logrus.WarnLevel
    }

    logrus.SetFormatter(&logrus.JSONFormatter{})
    logrus.SetFormatter(&logrus.TextFormatter{})
    logrus.SetLevel(logLevel)
}

func initGrpc() {
    logrus.Debugf("initGrpc() - START")

    listen, err := net.Listen("tcp", PORT)

    if err != nil {
        logrus.Fatalf("failed to listen: %v", err)
    }

    logrus.Debugf("gRPC listening in %s port", PORT)

    s := grpc.NewServer()
    pb.RegisterArduinoServer(s, &Controller{})
    reflection.Register(s)

    if err := s.Serve(listen); err != nil {
        logrus.Fatalf("failed to server: %v", err)
    }

    logrus.Debugf("initGrpc() - END")
}

func initDatabase() {
    logrus.Debugf("initDatabase() - START")

    if isDev() {
        time.Sleep(time.Second * 15)
    }

    db, err := gorm.Open("mysql", "iot:iot@tcp(iot_mysql:3306)/iot?charset=utf8&parseTime=True&loc=Local")
    db.LogMode(true)
    db.SetLogger(gormrus.New())

    if nil != err {
        logrus.Fatalf("Database connection error: %s", err)
    }

    logrus.Debug("Database connected")

    database = db

    logrus.Debugf("initDatabase() - END")
}

func stop() {
    defer database.Close()
    os.Exit(0)
}

func main() {
    if "" == ENV {
        ENV = "dev"
    }

    initLogger()
    initDatabase()
    initGrpc()

    stop()
}
