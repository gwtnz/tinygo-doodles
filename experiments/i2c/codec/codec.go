package main

import (
	"bytes"
	"fmt"
	"machine"
	"time"
)

// BufSz is the range of addresses to scan
const BufSz = 127

var output = false
var addrs [BufSz]uint8

func initI2CCodec() {
	I2C1Conf := machine.I2CConfig{Frequency: machine.TWI_FREQ_100KHZ}
	machine.AUDIO_DRIVER_NRESET_LINE.Configure(machine.PinConfig{Mode: machine.PinOutput})
	// Set NReset line low
	machine.AUDIO_DRIVER_NRESET_LINE.Set(false)
	machine.I2C1.Configure(I2C1Conf)
	// Give the CS43L22 some time to wake up
	time.Sleep(time.Millisecond * 10)
	// Set NReset line high
	machine.AUDIO_DRIVER_NRESET_LINE.Set(true)
}

func scanI2CBus() {
	w := []byte{0x01}
	r := []byte{0}

	var i uint16

	println("I2C bus scan")
	for i = 0x4a; i < 0x4b; i++ {
		err := machine.I2C1.Tx(i, w, r)
		if err == nil {
			addrs[i] = 1
			machine.UART0.WriteByte(0x2a)
		} else {
			println(err.Error())
			machine.I2C1.Reset()
			addrs[i] = 0
			machine.UART0.WriteByte(0x2e)
		}
		if (i+1)%16 == 0 {
			machine.UART0.WriteByte('\r')
			machine.UART0.WriteByte('\n')
		}
		time.Sleep(time.Millisecond * 2)
	}
	println()
}

func outputI2CBus() {
	// Wait till we have the semaphore
	for output {
		time.Sleep(time.Millisecond)
	}
	output = true

	var b bytes.Buffer
	b.WriteString("I2C scan bus\r\n")

	for i := 0; i < BufSz; i++ {
		if i%16 == 0 {
			b.WriteString(fmt.Sprintf("%02x ", i))
		}
		if addrs[i] > 0 {
			b.WriteString(fmt.Sprintf("%02x ", i))
		} else {
			b.WriteString(fmt.Sprintf(" . "))
		}
		if (i+1)%16 == 0 {
			b.WriteString("\r\n")
		}
	}
	println(b.String())

	output = false
}

func programID(msg string, timing time.Duration) {
	for {
		// Wait till we have the semaphore
		for output {
			time.Sleep(time.Millisecond * 51)
		}
		output = true
		println("\n" + msg)
		output = false
		time.Sleep(time.Millisecond * timing)
	}
}

func main() {
	output = false
	go programID("i2c codec", 5000)
	initI2CCodec()

	// loop for a while
	for {
		scanI2CBus()
		//outputI2CBus()
		time.Sleep(time.Second * 3)
	}
}
