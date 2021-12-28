package main

import "fmt"

var (
	actor   string  = "Mixko"
	geo     string  = "Geo"
	number1 int     = 1
	season  float32 = .3
)

const (
	yu string = "Hello"
	te string = "World"
	op int    = 50
)

func main() {
	fmt.Println("Hello")
	var i int
	i = 50
	k := 99 //Can't re assignment
	var t float32 = 5.5978

	var convert int
	convert = int(t)

	fmt.Println(convert)

	println(i, k, t)
	fmt.Printf("%v, %T", i, i)
	fmt.Println(actor)
	fmt.Println(yu, te, op)
	fmt.Println(actor, geo, number1, season)
}
