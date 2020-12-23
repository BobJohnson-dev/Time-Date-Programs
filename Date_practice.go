package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()

	  
    fmt.Println("Current Time in String: ", currentTime.String())
      
    fmt.Println("MM-DD-YYYY : ", currentTime.Format("01-02-2006"))
      
    fmt.Println("YYYY-MM-DD : ", currentTime.Format("2006-01-02"))
      
    fmt.Println("YYYY.MM.DD : ", currentTime.Format("2006.01.02 15:04:05"))
      
    fmt.Println("YYYY#MM#DD {Special Character} : ", currentTime.Format("2006#01#02"))
      
    fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2006-01-02 15:04:05"))

}

