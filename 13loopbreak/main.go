package main

import (
	"fmt"
)

func main() {

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday"}
	fmt.Println(days)

	// for d :=0; d<len(days);d++{
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }
	// for _, day := range days {
	// 	fmt.Printf("index is and value is %v\n", day)
	// }
	rougevalue := 1
	for rougevalue < 10 {

		if rougevalue == 2 {
			goto label1
		}
		if rougevalue == 5 {
			rougevalue++
			continue
		}
		fmt.Println("Value is", rougevalue)
		rougevalue++
	}
label1:
	fmt.Println("Welcome to GoLang")
}
