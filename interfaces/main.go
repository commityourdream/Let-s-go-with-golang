package main

import (
	"fmt"
	"log"
)

type calculation interface {
	Add() int64
	Multiple() float64
}

type meta interface {
	getLarge() int64
}

type numbers struct {
	x int64
	y int64
	numberlogy
}

type numberlogy struct {
	FilenName string
	FileSize  int64
	IsActive  bool
}

func (nml numberlogy) getLarge() int64{
	fmt.Sprintf("Files details are %s is size of %d which is active", nml.FilenName, nml.FileSize)
	return int64(0)
}

func (n numbers) Add() int64 {
	return n.x + n.y
}

func (n numbers) Multiple() float64 {
	return float64(n.x * n.y)
}

func emdeding(n meta) {
	fmt.Printf("check the")
}

func main() {
	num := numbers{
		x: 10,
		y: 20,
	}
	emdeding(num)
	fmt.Printf("reduce %d", num.x)
	DemoTypeAssertion(num)

}

func DemoTypeAssertion(n calculation) {
	num, ok := n.(numbers)
	if !ok {
		log.Printf("Failed to type assert")
		return
	}

	fmt.Printf("type assertion: %d", num.x)
	sum := n.Add()
	fmt.Printf("Sum of two numbers: %d\n", sum)
	mul := n.Multiple()
	fmt.Printf("Multiple of two numbers %f\n", mul)
}

func genericType(val interface{}) {
	intValue, ok := val.(int64)
	if !ok {
		log.Printf("Failed to get convert")
		return
	}
	fmt.Printf("Convert the value from interface to integer %d\n", intValue)
	fmt.Printf("Value of any type: %v\n", val)
}
