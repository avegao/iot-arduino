package main

import (
    "golang.org/x/net/context"
    pb "github.com/avegao/iot-arduino/proto"
    "github.com/avegao/iot-arduino/entity"
    "github.com/Sirupsen/logrus"
)

type Controller struct {}

func (controller *Controller) GetTemperature(ctx context.Context, in *pb.ArduinoRequest) (*pb.GetTemperatureResponse, error) {
    arduino := parseArduinoRequest(in)

    return &pb.GetTemperatureResponse{Temperature: arduino.GetArduinoTemp()}, nil
}

func (controller *Controller) IsPower(ctx context.Context, in *pb.ArduinoRequest) (*pb.PowerResponse, error) {
    arduino := parseArduinoRequest(in)

    return &pb.PowerResponse{Power: arduino.IsPower()}, nil
}

func (controller *Controller) PowerOn(ctx context.Context, in *pb.ArduinoRequest) (*pb.PowerResponse, error) {
    arduino := parseArduinoRequest(in)

    return &pb.PowerResponse{Power: arduino.PowerOn()}, nil
}

func (controller *Controller) PowerOff(ctx context.Context, in *pb.ArduinoRequest) (*pb.PowerResponse, error) {
    arduino := parseArduinoRequest(in)

    return &pb.PowerResponse{Power: arduino.PowerOff()}, nil
}

func (controller *Controller) ListAll(ctx context.Context, in *pb.ListAllRequest) (*pb.ListAllResponse, error) {
    arduinos := []entity.Arduino{}

    if err := database.Find(&arduinos).Error; nil != err {
        return nil, err
    }

    response := []*pb.ArduinoResponse{}

    for _, arduino := range arduinos {
        response = append(response, arduino.ToResponse())
    }

    return &pb.ListAllResponse{Results: response}, nil
}

func parseArduinoRequest(request *pb.ArduinoRequest) *ArduinoRequest {
    logrus.Debugf("parseArduinoRequest() - START")

    arduino := new(ArduinoRequest)
    arduino.Id = int32(request.Id)
    arduino.Name = request.Name
    arduino.Hostname = request.Url

    logrus.Debugf("parseArduinoRequest() - END")

    return arduino
}
