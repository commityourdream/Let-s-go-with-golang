package main

import "fmt"


func main()  {

	fmt.Println("Welcome to Functions in GoLang")
	greeter()
	result := adder(5,3)
	fmt.Println("sum of two number is:",result)

	proRes:= proadder(2,3,4,0,5,6,7)
	fmt.Println("Sum of n number is :",proRes)
	
}

func adder(val1 int, val2 int)int{
	return val1 + val2

}
func proadder(values ...int)int{
	total :=0
	for _,val := range values{
		total += val
	}
return total
}
func greeter(){
	fmt.Println("Namste from GoLang")

}