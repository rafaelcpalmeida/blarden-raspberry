package main

import (
    "fmt"
    "time"

    log "github.com/sirupsen/logrus"
    "github.com/stianeikeland/go-rpio/v4"
)

func recoverOpenDoorPanic() {
    if r := recover(); r != nil {
        log.Error("Recovered from error: ", r)
    }
}

func OpenDoor() {
    defer recoverOpenDoorPanic()
    err := rpio.Open()

    defer rpio.Close()
  
    if err != nil {
        panic(fmt.Sprint("unable to open gpio", err.Error()))
    }

    pin := rpio.Pin(24)
    pin.Output()

    pin.High()
    time.Sleep(1 * time.Second)
    pin.Low()
}
