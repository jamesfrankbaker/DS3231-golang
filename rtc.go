// i2c for Dallas semi. DS1307 and DS3231 clocks

package rtc

import "fmt"
import "strconv"
// import "time"
import "github.com/jamesfrankbaker/i2c"

var add uint8 = 0x68 // Clock i2c address goes here



func ReadDateTimeUint8() (uint8, uint8, uint8, uint8, uint8, uint8, uint8, uint8 ) {
// Reads time and date from DS1307 or 3231, returns numeric values

// Create new connection to I2C bus on 1 line with address add
  i2c, err := i2c.NewI2C(add, 1)
  if err != nil { fmt.Println("Error on i2c device open") }
  // Free I2C connection on exit
  defer i2c.Close()
if err != nil {
    panic(err)
}
var s uint8 // seconds from two BCD nibbles
var min uint8 // minutes from two BCD nibbles
var h uint8 // hours from two BCD nibbles
var w uint8 // day of week from two BCD nibbles
var d uint8 // date from two BCD nibbles
var mon uint8 // month from two BCD nibbles
var y uint8 // seconds from two BCD nibbles
var c uint8 // control register

t := []byte{0,0,0,0,0,0,0,0} // array/slice for DS1307 data

i2c.Write([]byte{0})

i2c.Read(t)

 // convert BCD to uint8, exclude non-numeric bits
  s=t[0]&0x0f + ((t[0]&0x70)>>4) * 10
  min=t[1]&0x0f + ((t[1]&0x70)>>4) * 10
  h=t[2]&0x0f + ((t[2]&0x30)>>4) * 10
  w=t[3]&0x07
  d=t[4]&0x0f + ((t[4]&0x30)>>4) * 10
  mon=t[5]&0x0f + ((t[5]&0x10)>>4) * 10
  y=t[6]&0x0f + ((t[6]&0xf0)>>4) * 10
  c=t[7]
 
  return s, min, h, w, d, mon, y, c
}




func ReadDateTimeString() (string) {
// reads DS1307 RTC, returns date & time string
// Uses readtd to get uint8 values from RTC, converts to string
var s uint8 // seconds in two BCD nibbles
var min uint8 // minutes in two BCD nibbles
var h uint8 // hours in two BCD nibbles
var w uint8 // dayofweek in two BCD nibbles
var d uint8 // date in two BCD nibbles
var mon uint8 // month in two BCD nibbles
var y uint8 // seconds in two BCD nibbles
var c uint8 // control register
var datime string  

s,min,h,w,d,mon,y,c = ReadDateTimeUint8()
  c = c&0x93 // to use c if not otherwise used, strips unused bits from control register
//  s=s+1 // to use s if not otherwise used
  w = w-1 // to use w if not otherwise used

datime = strconv.Itoa(int(d))+"/"+strconv.Itoa(int(mon))+"/"+strconv.Itoa(int(y))+" "+strconv.Itoa(int(h))+":"+strconv.Itoa(int(min))+":"+strconv.Itoa(int(s))
return datime
}



func WriteDateTimeUint8(s,min,h,w,d,mon,y,c uint8) {
// Writes time and date to DS1307
// This version does not write the control register "c"

var z uint8 = 0 // zero byte may need to be written first

// Create new connection to I2C bus on 1 line with address 0x68
  i2c, err := i2c.NewI2C(add, 1)
  if err != nil { fmt.Println("Error on i2c device open") }
  // Free I2C connection on exit
  defer i2c.Close()
if err != nil {
    panic(err)
}

// first convert decimal numbers to two nibble BCD
s = (s/10)<<4 + (s%10)
min = (min/10)<<4 + (min%10)
h = (h/10)<<4 + (h%10)
d = (d/10)<<4 + (d%10)
mon = (mon/10)<<4 + (mon%10)
y = (y/10)<<4 + (y%10)

i2c.Write([]byte{z,s,min,h,w,d,mon,y})

return	
}



