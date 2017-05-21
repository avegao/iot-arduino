package entity

import (
    "github.com/jinzhu/gorm"
    "github.com/avegao/iot-arduino/proto"
)

type Arduino struct {
    gorm.Model
    Name string
    URL string
    Token string
    Auto bool
}

func (arduino *Arduino) IsAuto() bool {
    return arduino.Auto
}

func (arduino *Arduino) ToResponse() *arduino_service.ArduinoResponse {
    response := new(arduino_service.ArduinoResponse)
    response.Id = uint32(arduino.ID)
    response.Name = arduino.Name
    response.Url = arduino.URL
    response.Auto = arduino.Auto

    return response
}
