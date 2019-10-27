package main

import (
    "fmt"
    "time"

    "github.com/stianeikeland/go-rpio/v4"
)

func OpenDoor() {
    err := rpio.Open()
  
    if err != nil {
        panic(fmt.Sprint("unable to open gpio", err.Error()))
    }
  
    pin := rpio.Pin(24)
    pin.Output()
  
    pin.High()
    time.Sleep(1 * time.Second)
    pin.Low()
  
    rpio.Close()
}
