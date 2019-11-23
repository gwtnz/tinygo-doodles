package main

// Use goroutines to spin off led timing for each of the user LEDs on the stm32f4 discovery board

import (
	"machine"
	"time"
)

func main() {
	var dur time.Duration = 1000
	go ledFn(machine.LED1, dur)
	dur <<= 1
	go ledFn(machine.LED2, dur)
	dur <<= 1
	go ledFn(machine.LED3, dur)
	dur <<= 1
	ledFn(machine.LED4, dur)
}

func ledFn(led machine.Pin, timing time.Duration) {
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
