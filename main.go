package main

import (
    "github.com/Sirupsen/logrus"
    "net"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "golang.org/x/net/context"
    pb "github.com/avegao/iotArduino/proto"
)

const (
    PORT = ":50000"
)

type server struct {}

func (s *server) GetTemperature(ctx context.Context, in *pb.ArduinoRequest) (*pb.GetTemperatureResponse, error) {
    arduino := parseArduinoRequest(in)

    return &pb.GetTemperatureResponse{Temperature: arduino.GetArduinoTemp()}, nil
}

func (s *server) IsPower(ctx context.Context, in *pb.ArduinoRequest) (*pb.PowerResponse, error) {
    arduino := parseArduinoRequest(in)

    return &pb.PowerResponse{Power: arduino.IsPower()}, nil
}

func (s *server) PowerOn(ctx context.Context, in *pb.ArduinoRequest) (*pb.PowerResponse, error) {
    arduino := parseArduinoRequest(in)

    return &pb.PowerResponse{Power: arduino.PowerOn()}, nil
}


func (s *server) PowerOff(ctx context.Context, in *pb.ArduinoRequest) (*pb.PowerResponse, error) {
    arduino := parseArduinoRequest(in)

    return &pb.PowerResponse{Power: arduino.PowerOff()}, nil
}


func init() {
    initLogger()
}

func initLogger() {
    logrus.SetFormatter(&logrus.JSONFormatter{})
    logrus.SetFormatter(&logrus.TextFormatter{})
    logrus.SetLevel(logrus.DebugLevel)
}

func main() {
    listen, err := net.Listen("tcp", PORT)

    if err != nil {
        logrus.Fatalf("failed to listen: %v", err)
    }

    logrus.Debugf("gRPC listening in %d port", PORT)

    s := grpc.NewServer()
    pb.RegisterArduinoServer(s, &server{})
    reflection.Register(s)

    if err := s.Serve(listen); err != nil {
        logrus.Fatalf("failed to server: %v", err)
    }
}

func parseArduinoRequest(request *pb.ArduinoRequest) *Arduino {
    arduino := new(Arduino)
    arduino.Id = request.Id
    arduino.Name = request.Name
    arduino.Hostname = request.Url

    return arduino
}
