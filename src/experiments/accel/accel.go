package main

// test the stm32f4 discovery board accelerometer, LIS302DL

import (
	"time"
	"tiny-go/tinygo/machine"
	"device/stm32"
)

// SPI pins
const (
	SPI0_SCK_PIN  = PA5
	SPI0_MOSI_PIN = PA7
	SPI0_MISO_PIN = PA6
)

// I2C pins
const (
	SDA_PIN = PA7
	SCL_PIN = PA5
)

// MEMs accelerometer
const (
  MEMS_ACCEL_CS = PE6
)

func spi1Init() {
	//code to configure PE3
RCC->AHB1ENR |= 1 << 4;	//enable clock to GPIOE
GPIOE->MODER |= 1 << 6;	//MODER3[1:0] = 01 bin
//code to enable AF - SPI1
RCC->AHB1ENR |= 1 << 0;	//enable clock to GPIOA
RCC->APB2ENR |= 1 << 12;	//clock to SPI1
GPIOA->AFR[0] |= (5<< 20);	//enable SPI CLK to PA5
GPIOA->AFR[0] |= (5<< 24);	//enable MISO to PA6
GPIOA->AFR[0] |= (5<< 28);	//enable MOSI to PA7
GPIOA->MODER &= ~(3 << 10);	//clear bits 10 & 11
GPIOA->MODER |= 2 << 10;	//MODER5[1:0] = 10 bin
GPIOA->MODER &= ~(3 << 12);	//clear bits 12 & 13
GPIOA->MODER |= 2 << 12;	//MODER6[1:0] = 10 bin
GPIOA->MODER &= ~(3 << 14);	//clear bits 14 & 15
GPIOA->MODER |= 2 << 14;	//MODER7[1:0] = 10 bin
SPI1->CR1	= 0x0003;	// CPOL=1, CPHA=1
SPI1->CR1	|= 1 << 2;	// Master Mode
SPI1->CR1	|= 1<<6;	// SPI enabled
SPI1->CR1	&= ~(7<<3);	// Use maximum frequency
SPI1->CR1	|= 3<<8;	// Soltware disables slave function
SPI1->CR2 = 0x0000;	//Motorola Format -
}

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
