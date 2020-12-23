package main
  
 import (
         "fmt"
         "time"    
 )
 
 func main() {       
    currentTime := time.Now()        
    fmt.Println("\n######################################\n")
    fmt.Println(currentTime.Format("2006-01-02 15:04:05"))
      
    fmt.Println("\n######################################\n")   
    timeStampString := currentTime.Format("2006-01-02 15:04:05")    
    layOut := "2006-01-02 15:04:05"    
    timeStamp, err := time.Parse(layOut, timeStampString)
    if err != nil {
        fmt.Println(err)          
    }   
    hr, min, sec := timeStamp.Clock()
      
    fmt.Println("Year   :", currentTime.Year())
    fmt.Println("Month  :", currentTime.Month())
    fmt.Println("Day    :", currentTime.Day())
    fmt.Println("Hour   :", hr)
    fmt.Println("Min    :", min)
    fmt.Println("Sec    :", sec)    
      
    fmt.Println("\n######################################\n")   
    year, month, day := time.Now().Date()
    fmt.Println("Year   :", year)
    fmt.Println("Month  :", month)
    fmt.Println("Day    :", day)
      
    fmt.Println("\n######################################\n")          
    t := time.Now()
      
    y := t.Year()
    mon := t.Month()
    d := t.Day()
    h := t.Hour()
    m := t.Minute()
    s := t.Second()
    n := t.Nanosecond()
      
    fmt.Println("Year   :",y)
    fmt.Println("Month   :",mon)
    fmt.Println("Day   :",d)
    fmt.Println("Hour   :",h)
    fmt.Println("Minute :",m)
    fmt.Println("Second :",s)
    fmt.Println("Nanosec:",n)
    }
    
