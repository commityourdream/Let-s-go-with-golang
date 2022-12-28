package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruitList = []string{"Apple", "Mango", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)
    
	fruitList = append(fruitList, "Tomatao","Banana")
	fmt.Println(fruitList)

	fruitList=append(fruitList[1:])
	fmt.Println(fruitList)

	highscores:=make([] int, 4)
	highscores[0]=543
	highscores[1]=100
	highscores[2]=200
	highscores[3]=300
	//highscores[4]=400
    highscores=append(highscores,333,876,120 )
	fmt.Println(highscores)

	sort.Ints(highscores)
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))
    
	// remove value from index

	var courses = [] string{"react","java","c","c++","swift","python"}

	var index int = 2
	courses= append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)

}
