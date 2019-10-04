package main

// This blinky is a bit more advanced than blink1, with two goroutines running
// at the same time and blinking a different LED. The delay of led2 is slightly
// less than half of led1, which would be hard to do without some sort of
// concurrency.

import (
	"machine"
	"time"
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
