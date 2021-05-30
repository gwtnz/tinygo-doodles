package main

type Servo interface {
	Init(min uint16, max uint16, freq uint32)
	Configure(pin Pin, pulseFreq uint8)
	Set(val uint8)
	Get() uint8
}

type PWM interface {
	Init(min uint32, max uint32, freq uint32)
	Configure(pin Pin)
	Set(val uint32)
	Get() uint32
}

