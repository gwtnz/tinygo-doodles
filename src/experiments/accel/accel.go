package main

// test the stm32f4 discovery board accelerometer, LIS302DL

import (
	"time"
	"tiny-go/tinygo/machine"
)

func main() {
	go ledFn(machine.LED1, 125)
	go ledFn(machine.LED2, 250)
	go ledFn(machine.LED3, 500)
	ledFn(machine.LED4, 1000)
}

func ledFn(led machine.Pin, timing time.Duration) {
	//led := machine.LED1
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		println("+")
		led.Low()
		time.Sleep(time.Millisecond * timing)

		println("-")
		led.High()
		time.Sleep(time.Millisecond * timing)
	}
}
