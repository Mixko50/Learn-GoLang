package main

import (
	"fmt"
	"unicode/utf8"
)

func main1() {
	println("Hello, World!")
	fmt.Printf("Hello %v\n", true)

	// Go has zero value (default value)
	var x int // 0
	var y int // 0
	y = x
	fmt.Println(x, y)

	// Array
	a := [3]int{}
	a[0] = 1

	// No limited array
	a2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	_ = a2
	fmt.Printf("%#v \n", a)

	// Slice (array without length) (list)
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s = append(s, 4)
	fmt.Printf("%#v\n%d\n", s, len(s))
	s2 := s[2:5]
	fmt.Printf("%#v\n", s2)

	// Each character is a byte, and each character has different length
	name := "Golang"
	name2 := "ก"
	fmt.Printf("%d\n%d\n%d\n", utf8.RuneCountInString(name), len(name2), utf8.RuneCountInString(name2))

	// Map (Dictionary)
	m := map[string]string{}
	m["name"] = "Golang"
	m["version"] = "1.18"

	testMap, ok := m["name"]
	if ok {
		println(testMap)
	} else {
		println("Not found")
	}

	// Tips
	t, u := 1, 2
	_, _ = t, u

	// Loop
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	for index, value := range s {
		println(index, value)
	}
}
