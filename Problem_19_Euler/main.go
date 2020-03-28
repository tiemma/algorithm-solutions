package main

import (
    "fmt"
    "flag"
    "time"
)



const layoutISO = "2006-01-02"


func main(){

    start := flag.String("start", "1901-01-01", "Start point in format yyyy-mm-dd")
    end := flag.String("end", "2000-12-31", "End point in format yyyy-mm-dd")
   
    start_date, _ := time.Parse(layoutISO, *start)
    end_date, _ := time.Parse(layoutISO, *end) 
    additional_date := 7 - int(start_date.Weekday())
    result := 0
    
    fmt.Println(additional_date)
    if additional_date < 7{
        start_date = start_date.Add(time.Hour * 24 * time.Duration(additional_date))
    }

    for start_date.Unix() < end_date.Unix() {
        if start_date.Day() == 1 {
            result += 1            
        }
        start_date = start_date.Add(time.Hour * 24 * 7)
    }
    fmt.Println(start_date.Day())
   //  result := (int(end_date.Sub(start_date).Hours() / 24) + additional_date) / 7
    
    
    fmt.Println(result)
}
