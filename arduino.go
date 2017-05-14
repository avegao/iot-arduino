package main

import (
    "fmt"
    log "github.com/Sirupsen/logrus"
)

var (
    getTemperatureUrl = "http://%s/arduino/temp"
    isPowerUrl = "http://%s/arduino/temp/power"
    powerOnUrl = "http://%s/arduino/temp/power/on"
    powerOffUrl = "http://%s/arduino/temp/power/off"
)

type Arduino struct {
    Id int32 `json:"id"`
    Name string `json:"name"`
    Hostname string `json:"hostname"`
}

type ArduinoTemperatureResponse struct {
    Temperature float32 `json:"temperature"`
}

type ArduinoPowerResponse struct {
    Power bool `json:"power"`
    Command string `json:"command"`
}

func (arduino Arduino) GetArduinoTemp() float32 {
    temperatureResponse := new(ArduinoTemperatureResponse)
    url := fmt.Sprintf(getTemperatureUrl, arduino.Hostname)

    err := GetJson(url, temperatureResponse)

    if nil != err {
        log.Error(err)
    }

    log.WithFields(log.Fields{
        "arduino": arduino.Name,
        "temperature": temperatureResponse.Temperature,
    }).Info("GetArduinoTemp")

    return temperatureResponse.Temperature
}

func (arduino Arduino) IsPower() bool {
    powerResponse := new(ArduinoPowerResponse)
    url := fmt.Sprintf(isPowerUrl, arduino.Hostname)

    err := GetJson(url, powerResponse)

    if nil != err {
        log.Error(err)
    }

    log.WithFields(log.Fields{
        "arduino": arduino.Name,
        "power": powerResponse.Power,
    }).Info("IsPower")

    return powerResponse.Power
}

func (arduino Arduino) PowerOn() bool {
    powerResponse := new(ArduinoPowerResponse)
    url := fmt.Sprintf(powerOnUrl, arduino.Hostname)

    err := GetJson(url, powerResponse)

    if nil != err {
        log.Error(err)
    }

    log.WithFields(log.Fields{
        "arduino": arduino.Name,
        "power": powerResponse.Power,
    }).Info("PowerOn")

    return powerResponse.Power
}

func (arduino Arduino) PowerOff() bool {
    powerResponse := new(ArduinoPowerResponse)
    url := fmt.Sprintf(powerOffUrl, arduino.Hostname)

    err := GetJson(url, powerResponse)

    if nil != err {
        log.Error(err)
    }

    log.WithFields(log.Fields{
        "arduino": arduino.Name,
        "power": powerResponse.Power,
    }).Info("PowerOff")

    return powerResponse.Power
}
