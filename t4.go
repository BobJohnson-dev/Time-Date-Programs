package main

import (
	"fmt"
	"time"
)

func main() { 
	
	//t4.go the same as t3.go
	//We’ll start by getting the current time.
	/* see https://gobyexample.com/time for this way of constructing a struc to print the time*/

	currentTime := time.Now()
    p := fmt.Println
	now := time.Now()
	p("At this time", now.Format("2006.01.02 15:04:05"))
	fmt.Println(currentTime.Format("2006.01.02 15:04:05"))
	

}
