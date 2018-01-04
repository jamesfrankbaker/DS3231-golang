// Based on Kidoman github embd example
// Adds rtc

package main

import (
	"github.com/jamesfrankbaker/rtc"
	"flag"
//	"time"
	"github.com/kidoman/embd"
	"github.com/kidoman/embd/controller/hd44780"
	_ "github.com/kidoman/embd/host/all"
)

func main() {
	
var datimold string
	
	flag.Parse()

	if err := embd.InitI2C(); err != nil {
		panic(err)
	}
	defer embd.CloseI2C()

	bus := embd.NewI2CBus(1)

	hd, err := hd44780.NewI2C(
		bus,
		// Use correct i2c address for your display
//		0x20,
		0x3F,
		hd44780.PCF8574PinMap,
		hd44780.RowAddress20Col,
		hd44780.TwoLine,
		hd44780.BlinkOn,
	)
	if err != nil {
		panic(err)
	}
	defer hd.Close()

	hd.Clear()
	
	hd.BacklightOn() // LCD is ready to go


/*	
	message := "Hello, world!"
	bytes := []byte(message)
	for _, b := range bytes {
		hd.WriteChar(b)
	}
	hd.SetCursor(0, 1)

	message = "@embd | hd44780"
	bytes = []byte(message)
	for _, b := range bytes {
		hd.WriteChar(b)
	}
	hd.SetCursor(0, 2)

	message = "@embd | hd44780line3"
	bytes = []byte(message)
	for _, b := range bytes {
		hd.WriteChar(b)
	}
	
*/	
for {
	datim := rtc.ReadDateTimeString()
	if datim != datimold {
		datimold = datim
	hd.SetCursor(0, 0)
	bytes := []byte(datim)
	for _, b := range bytes {
		hd.WriteChar(b)
	}
		hd.WriteChar(byte(32)) // space
		hd.WriteChar(byte(32)) // space
		hd.WriteChar(byte(32)) // space
		hd.BlinkOff() // 
	
	}	
	
//	time.Sleep(1 * time.Second)
//	hd.BacklightOff()
}
}
