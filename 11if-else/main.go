package main

import "fmt"

func main(){

	LoginCount:=10
	var result string
	if LoginCount<10{
		result="Regular User"
	}else if LoginCount>10{
		result="Watched Out"
	}else{
		result="Exactly result 10"
	}
	

	fmt.Println(result)
	if 9%2==0{
		fmt.Println("Even Number")
	}else{
		fmt.Println("Odd Number")
	}
	if num :=3; num <10{
		fmt.Println("Number is less than 10")
	}else{
		fmt.Println("Number is not less than 10")
	}
}