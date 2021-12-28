package main

import "fmt"

const myConst int = 12

const (
	_  = iota //ignore first value
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

const (
	_ = iota + 5
	ab
	cd
	ef
)

func main() {
	const myConst = 20
	var num int = 22
	fmt.Printf("%v, %T\n", myConst+num, myConst+num)

	var num2 int32 = 22
	//myConst is auto assigned to int32
	fmt.Printf("%v, %T\n", myConst+num2, myConst+num2)

	fmt.Println("----- iota -----")
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)

	fmt.Printf("%v\n", ab)
	var special int = 6
	fmt.Printf("%v\n", special == ab)
}
