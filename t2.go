package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
    fmt.Println(now.Month(), now.Day(), now.Year())
	fmt.Println("Month:", now.Month())
	fmt.Println("Day:", now.Day())
	fmt.Println("Hour:", now.Hour())
	fmt.Println("Minute:", now.Minute())
	fmt.Println("Second:", now.Second())
	fmt.Println("Time", now)
	fmt.Println("<<<<<<<<<>>>>>>>>>")
	fmt.Println("Time: ", now.Format("15:04:05"))
	fmt.Println("Date:", now.Format("Jan 2, 2006"))
	fmt.Println("Timestamp:", now.Format(time.Stamp))
	fmt.Println("ANSIC:", now.Format(time.ANSIC))
	fmt.Println("UnixDate:", now.Format(time.UnixDate))
	fmt.Println("Kitchen:", now.Format(time.Kitchen))
}
