package main

import "fmt"


func main(){

	var FruitList [4]string
    FruitList[0]="Apple"
	FruitList[1]="Grapes"
	FruitList[2]="Peach"
	FruitList[3]="Tomato"
	fmt.Println("Fruitlist is: ",FruitList)
	fmt.Println("Fruitlist is: ",len(FruitList))

	var vegList = [5]string{"potato","beans","mushroom"}
	fmt.Println("Veglist is: ",vegList)
	fmt.Println("Veglist is: ",len(vegList))
}