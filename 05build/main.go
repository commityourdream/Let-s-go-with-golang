package main

import (
	"fmt"
	"time"
)

func main() {
	PresentTime := time.Now()
	fmt.Println("Current time is:", PresentTime)
	fmt.Println(PresentTime.Format("01-02-2006 Monday"))
	createDate := time.Date(2022, time.December, 28, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}
